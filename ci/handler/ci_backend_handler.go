package handler

import (
	"ci-backend/dao"
	"ci-backend/thrift/ci"
	"context"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type CIBackendServiceHandler struct{}

func (h *CIBackendServiceHandler) GetRepoNames(
	ctx context.Context,
) (r []string, err error) {
	res := []string{}
	for repoName := range dao.WatchedRepos {
		res = append(res, repoName)
	}
	return res, nil
}

func (h *CIBackendServiceHandler) GetRepoInfo(
	ctx context.Context,
	name string,
) (r *ci.RepoInfo, err error) {
	// build result object
	res := ci.RepoInfo{}
	res.Name = name
	res.MaxRunNum = 0 // TODO: temp solution, fix this

	// find the repo
	repo, ok := dao.WatchedRepos[name]
	if !ok {
		return nil, nil
	}

	// get branches info
	branches, err := repo.Branches()
	if err != nil {
		return nil, nil
	}
	err = branches.ForEach(func(br *plumbing.Reference) error {
		res.BranchNames = append(res.BranchNames, br.Name().String())
		return nil
	})
	if err != nil {
		return nil, nil
	}

	// get commits info
	commits, err := repo.CommitObjects()
	if err != nil {
		return nil, nil
	}
	err = commits.ForEach(func(c *object.Commit) error {
		res.CommitHashs = append(res.CommitHashs, c.Hash.String())
		return nil
	})
	if err != nil {
		return nil, nil
	}

	return &res, nil
}

func (h *CIBackendServiceHandler) GetBranchInfo(
	ctx context.Context,
	repoName string,
	branchName string,
) (r *ci.BranchInfo, err error) {
	// build result object
	res := ci.BranchInfo{}
	res.Name = branchName
	res.RunNums = make([]int32, 0) // TODO: temp solution, fix this

	// find the repo
	repo, ok := dao.WatchedRepos[repoName]
	if !ok {
		return nil, nil
	}

	// get commit hash of the branch
	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branchName), true)
	if ref == nil || err != nil {
		return nil, nil
	}
	res.CommitHash = ref.Hash().String()
	return &res, nil
}

func (h *CIBackendServiceHandler) GetCommitInfo(
	ctx context.Context,
	repoName string,
	commitHash string,
) (r *ci.CommitInfo, err error) {
	// build result object
	res := ci.CommitInfo{}
	res.Hash = commitHash
	res.RunNums = make([]int32, 0) // TODO: temp solution, fix this

	// find the repo
	repo, ok := dao.WatchedRepos[repoName]
	if !ok {
		return nil, nil
	}

	// get commit hash of the branch
	ref, err := repo.CommitObject(plumbing.NewHash(commitHash))
	if ref == nil || err != nil {
		return nil, nil
	}
	res.Msg = ref.Message
	res.Author = ref.Author.Name
	return &res, nil
}

func (h *CIBackendServiceHandler) GetRunInfo(
	ctx context.Context,
	repoName string,
	runNum int32,
) (r *ci.RunInfo, err error) {
	// build result object
	res := ci.RunInfo{}
	res.Num = runNum

	// find the repo
	_, ok := dao.WatchedRepos[repoName]
	if !ok {
		return nil, nil
	}

	// TODO: fetch run info from db
	return nil, nil
}
