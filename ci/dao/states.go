package dao

import "github.com/go-git/go-git/v5"

var WatchedRepos map[string]git.Repository

func init() {
	WatchedRepos = make(map[string]git.Repository)
}
