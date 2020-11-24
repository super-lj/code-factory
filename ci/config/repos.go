package config

import (
	"github.com/go-git/go-git/v5"
)

var WatchedRepos map[string]*git.Repository

func InitWatchedRepos() error {
	WatchedRepos = make(map[string]*git.Repository)
	for repoName, config := range CurrConfig.Repos {
		repo, err := git.PlainOpen(config.Path)
		if err != nil {
			return err
		}
		WatchedRepos[repoName] = repo
	}
	return nil
}
