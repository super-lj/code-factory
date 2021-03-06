package handler

import (
	"ci-backend/config"
	"ci-backend/dao"
	"ci-backend/thrift/ci"
	"context"
	"sort"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type CIBackendServiceHandler struct{}

func (h *CIBackendServiceHandler) GetRepoNames(
	ctx context.Context,
) (r []string, err error) {
	res := []string{}
	for repoName := range config.WatchedRepos {
		res = append(res, repoName)
	}
	sort.Strings(res)
	return res, nil
}

func (h *CIBackendServiceHandler) GetRepoInfo(
	ctx context.Context,
	name string,
) (r *ci.RepoInfo, err error) {
	// build result object
	res := ci.RepoInfo{}
	res.Name = name

	// find the repo
	repo, ok := config.WatchedRepos[name]
	if !ok {
		return nil, nil
	}

	// get branches info
	branches, err := repo.Branches()
	if err != nil {
		return nil, nil
	}
	err = branches.ForEach(func(br *plumbing.Reference) error {
		res.BranchNames = append(res.BranchNames, br.Name().Short())
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

	// get max num of a repo from DB
	rows, err := dao.RunDB.
		Model(&dao.Run{}).
		Select("max(num) as max_num").
		Group("repo_name").
		Having("repo_name = ?", name).
		Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&res.MaxRunNum)
		if err != nil {
			return nil, err
		}
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

	// find the repo
	repo, ok := config.WatchedRepos[repoName]
	if !ok {
		return nil, nil
	}

	// get commit hash of the branch
	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branchName), true)
	if ref == nil || err != nil {
		return nil, nil
	}
	res.CommitHash = ref.Hash().String()

	// get all runs of the branch from DB
	runs := make([]dao.Run, 0)
	dbErr := dao.RunDB.
		Where("repo_name = ? AND branch_name = ?", repoName, branchName).
		Find(&runs).
		Error
	if dbErr == nil {
		for _, r := range runs {
			res.RunNums = append(res.RunNums, r.Num)
		}
	}

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

	// find the repo
	repo, ok := config.WatchedRepos[repoName]
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

	// get all runs of the branch from DB
	runs := make([]dao.Run, 0)
	dbErr := dao.RunDB.
		Where("repo_name = ? AND commit_hash = ?", repoName, commitHash).
		Find(&runs).
		Error
	if dbErr == nil {
		for _, r := range runs {
			res.RunNums = append(res.RunNums, r.Num)
		}
	}

	return &res, nil
}

func (h *CIBackendServiceHandler) GetRunInfo(
	ctx context.Context,
	repoName string,
	runNum int32,
) (r *ci.RunInfo, err error) {
	run := dao.Run{}
	dbErr := dao.RunDB.
		Where("repo_name = ? AND num = ?", repoName, runNum).
		Find(&run).
		Error
	if dbErr != nil {
		return nil, nil
	}

	res := ci.RunInfo{}
	res.Num = runNum
	res.StartTimestamp = int32(run.CreatedAt.Unix())
	if run.Status == "IN_PROGRESS" {
		res.Duration = int32(time.Now().Sub(run.CreatedAt).Seconds())
	} else {
		res.Duration = int32(run.UpdatedAt.Sub(run.CreatedAt).Seconds())
	}
	res.Status = run.Status
	res.Log = run.Log
	res.BranchName = run.BranchName
	res.CommitHash = run.CommitHash
	return &res, nil
}
