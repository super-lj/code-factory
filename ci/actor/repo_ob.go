package actor

import (
	"ci-backend/config"
	"ci-backend/dao"
	"errors"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func StartRepoObserver(repoName string) (chan string, error) {
	repo, ok := config.WatchedRepos[repoName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Cannot find repo %s ", repoName))
	}
	ch := make(chan string)
	go RepoObserverMainLoop(repoName, repo, ch)
	return ch, nil
}

func RepoObserverMainLoop(repoName string, repo *git.Repository, ch chan string) {
	// store the hashes of all commmits at init
	commIter, err := repo.CommitObjects()
	if err != nil {
		ch <- fmt.Sprintf("Cannot get commits set of repo [%s]: %v", repoName, err)
		return
	}
	prevHashes := mapset.NewSet()
	commIter.ForEach(func(c *object.Commit) error {
		prevHashes.Add(c.Hash.String())
		return nil
	})

	// get the largest run num from DB
	var nextRunNum int32 = 0
	rows, err := dao.RunDB.
		Model(&dao.Run{}).
		Select("max(num) as max_num").
		Group("repo_name").
		Having("repo_name = ?", repoName).
		Rows()
	if err != nil {
		ch <- fmt.Sprintf("Cannot read max num of repo [%s] from DB: %v", repoName, err)
		return
	}
	for rows.Next() {
		err := rows.Scan(&nextRunNum)
		if err != nil {
			ch <- fmt.Sprintf("Cannot read max num of repo [%s] from DB: %v", repoName, err)
			return
		}
	}
	nextRunNum++

	// init array of ch from pipeline executors
	execChs := []chan string{}

	// start the loop
	for {
		// get all the hashes of commits
		commIter, err := repo.CommitObjects()
		if err != nil {
			ch <- fmt.Sprintf("Cannot get commits set of repo [%s]: %v", repoName, err)
			return
		}
		hashes := mapset.NewSet()
		commIter.ForEach(func(c *object.Commit) error {
			hashes.Add(c.Hash.String())
			return nil
		})

		// check if there is new commit
		if hashes.IsProperSuperset(prevHashes) {
			// get diff of commits
			diff := hashes.Difference(prevHashes).ToSlice()

			// get the name of current branch
			ref, err := repo.Head()
			if err != nil {
				ch <- fmt.Sprintf("Cannot read current branch name of repo [%s]: %v", repoName, err)
				return
			}

			// start new pipeline execution
			for _, h := range diff {
				exeCh, err := StartPipelineExecutor(repoName, ref.Name().Short(), h.(string), nextRunNum)
				if err != nil {
					ch <- fmt.Sprintf("Cannot start pipeline exec : %v", err)
					continue
				}
				execChs = append(execChs, exeCh)
				nextRunNum++
			}
			prevHashes = hashes
		}

		// check if there is message from executors
		for _, c := range execChs {
			select {
			case msg := <-c:
				ch <- msg
			default:
				continue
			}
		}

		// check if the loop needs to be stopped
		select {
		case msg := <-ch:
			if msg == "stop" {
				break
			}
		default:
		}

		// sleep for a while to give scheduler a chance to schedule
		time.Sleep(300 * time.Millisecond)
	}
}
