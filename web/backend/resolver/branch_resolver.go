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

type BranchResolver struct {
	id       graphql.ID
	name     string
	repoName string
}

func (r *BranchResolver) Id() graphql.ID {
	return r.id
}

func (r *BranchResolver) Name() string {
	return r.name
}

func (r *BranchResolver) Commit(ctx context.Context) (*CommitResolver, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.name)
	branchInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["branch_info"]
	res, err := branchInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	brInfo := res.(*ci.BranchInfo)
	if brInfo == nil {
		return nil, nil
	}

	// build commit resolver
	cmRx := &CommitResolver{
		id:       graphql.ID(r.repoName + " " + brInfo.CommitHash),
		repoName: r.repoName,
		hash:     brInfo.CommitHash,
	}
	return cmRx, nil
}

func (r *BranchResolver) RunsConnection(
	ctx context.Context,
	args struct {
		First *int32
		After *graphql.ID
	}) (*BranchRunsConnectionResolver, error) {
	// load branch info
	fetchKey := fmt.Sprintf("%s,%s", r.repoName, r.name)
	branchInfoloader := ctx.Value("loaders").(map[string]*dataloader.Loader)["branch_info"]
	res, err := branchInfoloader.Load(ctx, dataloader.StringKey(fetchKey))()
	if err != nil {
		return nil, err
	}
	brInfo := res.(*ci.BranchInfo)
	if brInfo == nil {
		return nil, nil
	}

	// sort it to be large num first
	sort.Slice(brInfo.RunNums, func(i, j int) bool {
		return brInfo.RunNums[i] > brInfo.RunNums[j]
	})

	// calculate the start and end
	start := 0
	if args.After != nil {
		for ; start < len(brInfo.RunNums); start++ {
			runID := graphql.ID(r.repoName + " " + strconv.Itoa(int(brInfo.RunNums[start])))
			if runID == *args.After {
				start++
				break
			}
		}
	}
	end := len(brInfo.RunNums)
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + int(*args.First)
	}
	if end > len(brInfo.RunNums) {
		end = len(brInfo.RunNums)
	}

	// build next level resolver
	brRunRx := &BranchRunsConnectionResolver{
		pageInfo: &PageInfoResolver{end != len(brInfo.RunNums)},
	}
	if start >= end {
		return brRunRx, nil
	}
	for _, num := range brInfo.RunNums[start:end] {
		// build and append edge resolver
		id := graphql.ID(r.repoName + " " + strconv.Itoa(int(num)))
		brRunRx.edges = append(brRunRx.edges, &BranchRunsEdgeResolver{
			cursor: id,
			node: &RunResolver{
				id:       id,
				repoName: r.repoName,
				num:      num,
			},
		})
	}
	return brRunRx, nil
}
