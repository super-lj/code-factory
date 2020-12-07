package resolver

import (
	"context"
	"fmt"
	"strconv"
	"web-backend/thrift/ci"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type RunResolver struct {
	id       graphql.ID
	repoName string
	num      int32
}

func (r *RunResolver) Id() graphql.ID {
	return r.id
}

func (r *RunResolver) Num() int32 {
	return r.num
}

func (r *RunResolver) StartTimestamp(ctx context.Context) (int32, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return 0, err
	}
	return res.(*ci.RunInfo).StartTimestamp, nil
}

func (r *RunResolver) Duration(ctx context.Context) (int32, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return 0, err
	}
	return res.(*ci.RunInfo).Duration, nil
}

func (r *RunResolver) Status(ctx context.Context) (string, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return "", err
	}
	return res.(*ci.RunInfo).Status, nil
}

func (r *RunResolver) Log(ctx context.Context) (string, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return "", err
	}
	return res.(*ci.RunInfo).Log, nil
}

func (r *RunResolver) Branch(ctx context.Context) (*BranchResolver, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	runInfo := res.(*ci.RunInfo)
	if runInfo == nil {
		return nil, nil
	}
	brID := graphql.ID(r.repoName + " " + runInfo.BranchName)
	brRx := &BranchResolver{
		id:       brID,
		name:     runInfo.BranchName,
		repoName: r.repoName,
	}
	return brRx, nil
}

func (r *RunResolver) Commit(ctx context.Context) (*CommitResolver, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, strconv.Itoa(int(r.num)))
	runInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["run_info"]
	res, err := runInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	runInfo := res.(*ci.RunInfo)
	if runInfo == nil {
		return nil, nil
	}
	cmID := graphql.ID(r.repoName + " " + runInfo.CommitHash)
	cmRx := &CommitResolver{
		id:       cmID,
		repoName: r.repoName,
		hash:     runInfo.CommitHash,
	}
	return cmRx, nil
}
