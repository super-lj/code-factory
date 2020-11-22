package actor

import (
	"ci-backend/dao"
	"errors"
	"fmt"
	"log"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func StartRepoObserver(repoName string) (chan string, error) {
	repo, ok := dao.WatchedRepos[repoName]
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
			diff := hashes.Difference(prevHashes).ToSlice()
			log.Printf("diff: %v\n", diff)
			// TODO: start new pipeline execution
			prevHashes = hashes
		}

		// check if the loop needs to be stopped
		select {
		case msg := <-ch:
			if msg == "stop" {
				break
			}
		default:
			continue
		}
	}
}
