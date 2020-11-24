package loader

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/graph-gophers/dataloader"
)

// batch functions
var getRepoNameBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	cli, tp, err := CreateCIBackendClient(CIBackendAddr)
	if err != nil {
		log.Print(err)
		return results
	}
	defer tp.Close()

	for i := 0; i < len(keys.Keys()); i++ {
		repoNames, err := cli.GetRepoNames(context.Background())
		if err != nil {
			log.Print(err)
			continue
		}
		results = append(results, &dataloader.Result{Data: repoNames})
	}
	return results
}

var getRepoInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	cli, tp, err := CreateCIBackendClient(CIBackendAddr)
	if err != nil {
		log.Print(err)
		return results
	}
	defer tp.Close()

	for _, name := range keys.Keys() {
		repoInfo, err := cli.GetRepoInfo(context.Background(), name)
		if err != nil {
			log.Print(err)
			continue
		}
		results = append(results, &dataloader.Result{Data: repoInfo})
	}
	return results
}

var getBranchInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	cli, tp, err := CreateCIBackendClient(CIBackendAddr)
	if err != nil {
		log.Print(err)
		return results
	}
	defer tp.Close()

	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		branchName := strs[1]
		brInfo, err := cli.GetBranchInfo(context.Background(), repoName, branchName)
		if err != nil {
			log.Print(err)
			continue
		}
		results = append(results, &dataloader.Result{Data: brInfo})
	}
	return results
}

var getCommitInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	cli, tp, err := CreateCIBackendClient(CIBackendAddr)
	if err != nil {
		log.Print(err)
		return results
	}
	defer tp.Close()

	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		commitHash := strs[1]
		cmInfo, err := cli.GetCommitInfo(context.Background(), repoName, commitHash)
		if err != nil {
			log.Print(err)
			continue
		}
		results = append(results, &dataloader.Result{Data: cmInfo})
	}
	return results
}

var getRunInfoBatchFn = func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var results []*dataloader.Result

	cli, tp, err := CreateCIBackendClient(CIBackendAddr)
	if err != nil {
		log.Print(err)
		return results
	}
	defer tp.Close()

	for _, k := range keys.Keys() {
		strs := strings.Split(k, ",")
		if len(strs) != 2 {
			return []*dataloader.Result{}
		}
		repoName := strs[0]
		runNum, err := strconv.ParseInt(strs[1], 10, 32)
		if err != nil {
			log.Print(err)
			continue
		}
		runInfo, err := cli.GetRunInfo(context.Background(), repoName, int32(runNum))
		if err != nil {
			log.Print(err)
			continue
		}
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
