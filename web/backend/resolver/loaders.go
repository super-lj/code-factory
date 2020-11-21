package resolver

import (
	"context"
	"strconv"
	"strings"
	"web-backend/mock"

	"github.com/graph-gophers/dataloader"
)

// batch functions
var getRepoNameBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for i := 0; i < len(keys.Keys()); i++ {
		results = append(results, &dataloader.Result{Data: mock.GetRepoNames()})
	}
	return results
}

var getRepoInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, name := range keys.Keys() {
		results = append(results, &dataloader.Result{Data: mock.GetRepoInfo(name)})
	}
	return results
}

var getBranchInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		branchName := strs[1]
		brInfo := mock.GetBranchInfo(repoName, branchName)
		results = append(results, &dataloader.Result{Data: brInfo})
	}
	return results
}

var getCommitInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		commitHash := strs[1]
		commmitInfo := mock.GetCommitInfo(repoName, commitHash)
		results = append(results, &dataloader.Result{Data: commmitInfo})
	}
	return results
}

var getRunInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result
	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		runNum, err := strconv.ParseInt(strs[1], 10, 32)
		if err != nil {
			return []*dataloader.Result{}
		}
		runInfo := mock.GetRunInfo(repoName, int32(runNum))
		results = append(results, &dataloader.Result{Data: runInfo})
	}
	return results
}

// Loaders
var RepoNameloader = dataloader.NewBatchedLoader(getRepoNameBatchFn)
var RepoInfoloader = dataloader.NewBatchedLoader(getRepoInfoBatchFn)
var BranchInfoloader = dataloader.NewBatchedLoader(getBranchInfoBatchFn)
var CommitInfoloader = dataloader.NewBatchedLoader(getCommitInfoBatchFn)
var RunInfoloader = dataloader.NewBatchedLoader(getRunInfoBatchFn)
