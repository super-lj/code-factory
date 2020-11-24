package image

import (
	"ci-backend/config"
	"context"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func InitDockerImages() error {
	for _, config := range config.CurrConfig.Images {
		err := BuildImage(config.Path, config.Tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func BuildImage(path string, tag string) error {
	// get docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	// check if the image is already included
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}
	for _, img := range images {
		if len(img.RepoTags) > 0 && img.RepoTags[0] == tag {
			log.Printf("docker image [%s] already exists!", img.RepoTags[0])
			return nil
		}
	}

	// image doesn't exist, build it
	buildCtx, err := archive.Tar(path, archive.Uncompressed)
	resp, err := cli.ImageBuild(ctx, buildCtx, types.ImageBuildOptions{
		Tags: []string{tag},
	})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
	return nil
}
