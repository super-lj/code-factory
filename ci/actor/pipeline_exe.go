package actor

import (
	"archive/tar"
	"bufio"
	"ci-backend/config"
	"ci-backend/dao"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/shlex"
	"gopkg.in/yaml.v2"
)

type Job struct {
	Name string
	Run  string
}

type CIConfig struct {
	Jobs []Job
}

func StartPipelineExecutor(
	repoName string,
	branchName string,
	commitHash string,
	runNum int32,
) (chan string, error) {
	config, ok := config.CurrConfig.Repos[repoName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Cannot find repo %s ", repoName))
	}
	ch := make(chan string)
	go PipelineExecMain(repoName, branchName, commitHash, runNum, ch, config.Path)
	return ch, nil
}

func PipelineExecMain(
	repoName string,
	branchName string,
	commitHash string,
	runNum int32,
	ch chan string,
	repoPath string,
) {
	log.Printf(
		"Started pipeline exec of repo [%s], branch [%s], commit [%s], num [%d]",
		repoName, branchName, commitHash, runNum)

	// get docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		ch <- fmt.Sprintf("Failed to init docker client: %v", err)
		return
	}

	// create & start container
	imgName := config.CurrConfig.Repos[repoName].Img
	image, ok := config.CurrConfig.Images[imgName]
	if !ok {
		ch <- fmt.Sprintf("Cannot find image: %s", imgName)
		return
	}

	containerCreateResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:     image.Tag,
		Tty:       true,
		OpenStdin: true,
	}, nil, nil, "")
	if err != nil {
		ch <- err.Error()
		return
	}

	containerID := containerCreateResp.ID
	if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		ch <- fmt.Sprintf("Failed to start ci container: %v", err)
		return
	}

	defer func() {
		// stop the container when exit
		if err := cli.ContainerStop(ctx, containerID, nil); err != nil {
			ch <- fmt.Sprintf("Failed to stop container: %v", err)
		}
	}()

	// copy repo to container
	// use `docker cp`in cli
	err = exec.
		Command("docker", "cp", repoPath, fmt.Sprintf("%s:/root/", containerID)).
		Run()
	if err != nil {
		ch <- fmt.Sprintf("Failed to cp repo into container: %v", err)
		return
	}

	// checkout specific commit in container
	execCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		WorkingDir: path.Join("/root/", path.Base(repoPath)),
		Cmd:        []string{"git", "checkout", commitHash},
	})
	if err != nil {
		ch <- fmt.Sprintf("Failed to create git checkout exec: %v", err)
		return
	}

	err = cli.ContainerExecStart(ctx, execCreateResp.ID, types.ExecStartCheck{})
	if err != nil {
		ch <- fmt.Sprintf("Failed to run git checkout: %v", err)
		return
	}

	// read yml
	r, _, err := cli.CopyFromContainer(
		ctx,
		containerID,
		path.Join("/root/", path.Base(repoPath), ".codefactory.yml"),
	)
	if err != nil {
		ch <- fmt.Sprintf("Failed to read .codefactory.yml: %v", err)
		return
	}

	tr := tar.NewReader(r)
	_, err = tr.Next()
	if err != nil {
		ch <- fmt.Sprintf("Failed to read .codefactory.yml: %v", err)
		return
	}

	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		ch <- fmt.Sprintf("Failed to read .codefactory.yml: %v", err)
		return
	}

	ciConfig := CIConfig{}
	err = yaml.Unmarshal(buf, &ciConfig)
	if err != nil {
		ch <- fmt.Sprintf("Failed to unmarshal .codefactory.yml: %v", err)
		return
	}

	// init run log in memory
	logSB := strings.Builder{}

	// store run info in DB
	runInfo := dao.Run{
		RepoName:   repoName,
		BranchName: branchName,
		CommitHash: commitHash,
		Num:        runNum,
		Status:     "IN_PROGRESS",
		Log:        logSB.String(),
	}
	err = dao.RunDB.Create(&runInfo).Error
	if err != nil {
		ch <- fmt.Sprintf("Cannot create new run record in DB: %v", err)
		return
	}

	// execute commands in yml and store logs
	execSucceed := true
	for _, job := range ciConfig.Jobs {
		// split run command into tokens
		tokens, err := shlex.Split(job.Run)
		if err != nil {
			ch <- fmt.Sprintf("Failed to split command [%s] exec: %v", job.Run, err)
			return
		}

		// create exec
		execCreateResp, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
			WorkingDir:   path.Join("/root/", path.Base(repoPath)),
			Cmd:          tokens,
			Tty:          true,
			AttachStderr: true,
			AttachStdout: true,
			AttachStdin:  true,
			Detach:       true,
		})
		if err != nil {
			ch <- fmt.Sprintf("Failed to create command [%s] exec: %v", job.Run, err)
			return
		}

		// attach and run exec
		attach, err := cli.ContainerExecAttach(ctx, execCreateResp.ID, types.ExecStartCheck{})
		if err != nil {
			ch <- fmt.Sprintf("Failed to attach command [%s] exec: %v", job.Run, err)
			return
		}
		defer attach.Close()

		outReader, outWriter := io.Pipe()
		outDone := make(chan error)
		go func() {
			_, err = stdcopy.StdCopy(outWriter, outWriter, attach.Reader)
			outWriter.Close()
			outDone <- err
		}()

		scanner := bufio.NewScanner(outReader)
		for scanner.Scan() {
			// append log to string builder
			logSB.WriteString(scanner.Text())
			logSB.WriteByte('\n')

			// update log of run info in DB
			err = dao.RunDB.Model(&runInfo).Update("log", logSB.String()).Error
			if err != nil {
				ch <- fmt.Sprintf("Failed to update log to DB: %v", err)
				return
			}
		}
		if err := scanner.Err(); err != nil {
			ch <- fmt.Sprintf("Failed to read from exec: %v", err)
			return
		}
		if err := <-outDone; err != nil {
			ch <- fmt.Sprintf("Stdcopy failed: %v", err)
			return
		}

		// get the exit code and determine if execution succeed
		inspRes, err := cli.ContainerExecInspect(ctx, execCreateResp.ID)
		if err != nil {
			ch <- fmt.Sprintf("Failed to read exec status: %v", err)
			return
		}
		if inspRes.ExitCode != 0 {
			execSucceed = false
			break
		}
	}

	// update status of run info in DB
	statusStr := "IN_PROGRESS"
	if execSucceed {
		statusStr = "SUCCEED"
	} else {
		statusStr = "FAILED"
	}
	err = dao.RunDB.Model(&runInfo).Update("status", statusStr).Error
	if err != nil {
		ch <- fmt.Sprintf("Failed to update status to DB: %v", err)
		return
	}

	log.Printf(
		"Pipeline exec of repo [%s], branch [%s], commit [%s], num [%d] completed: %s",
		repoName, branchName, commitHash, runNum, statusStr)
}
