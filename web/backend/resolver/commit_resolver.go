package resolver

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"web-backend/thrift/ci"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type CommitResolver struct {
	id       graphql.ID
	repoName string
	hash     string
}

func (r *CommitResolver) Id() graphql.ID {
	return r.id
}

func (r *CommitResolver) Hash() string {
	return r.hash
}

func (r *CommitResolver) Msg(ctx context.Context) (string, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.hash)
	commitInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["commit_info"]
	res, err := commitInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return "", err
	}
	return res.(*ci.CommitInfo).Msg, nil
}

func (r *CommitResolver) Author(ctx context.Context) (string, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.hash)
	commitInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["commit_info"]
	res, err := commitInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return "", err
	}
	return res.(*ci.CommitInfo).Author, nil
}

func (r *CommitResolver) RunsConnection(
	ctx context.Context,
	args struct {
		First *int32
		After *graphql.ID
	}) (*CommitRunsConnectionResolver, error) {
	// fetch commit info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.hash)
	commitInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["commit_info"]
	res, err := commitInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	if res.(*ci.CommitInfo) == nil {
		return nil, nil
	}
	commitInfo := res.(*ci.CommitInfo)

	// sort it to be large num first
	sort.Slice(commitInfo.RunNums, func(i, j int) bool {
		return commitInfo.RunNums[i] > commitInfo.RunNums[j]
	})

	// calculate the start and end
	start := 0
	if args.After != nil {
		for ; start < len(commitInfo.RunNums); start++ {
			runID := graphql.ID(r.repoName + " " + strconv.Itoa(int(commitInfo.RunNums[start])))
			if runID == *args.After {
				start++
				break
			}
		}
	}
	end := len(commitInfo.RunNums)
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + int(*args.First)
	}
	if end > len(commitInfo.RunNums) {
		end = len(commitInfo.RunNums)
	}

	// build next level resolver
	cmRunsRx := &CommitRunsConnectionResolver{
		pageInfo: &PageInfoResolver{end != len(commitInfo.RunNums)},
	}
	if start >= end {
		return cmRunsRx, nil
	}
	for _, num := range commitInfo.RunNums[start:end] {
		runId := graphql.ID(r.repoName + " " + strconv.Itoa(int(num)))
		cmRunsRx.edges = append(cmRunsRx.edges, &CommitRunsEdgeResolver{
			cursor: runId,
			node: &RunResolver{
				id:       runId,
				repoName: r.repoName,
				num:      num,
			},
		})
	}
	return cmRunsRx, nil
}
