package resolver

import (
	"context"
	"strconv"
	"web-backend/loader"
	"web-backend/thrift/ci"

	"github.com/graph-gophers/dataloader"
	graphql "github.com/graph-gophers/graphql-go"
)

type RepoResolver struct {
	id   graphql.ID
	name string
}

func (r *RepoResolver) Id() graphql.ID {
	return r.id
}

func (r *RepoResolver) Name() string {
	return r.name
}

func (r *RepoResolver) BranchesConnection(args struct {
	First *int32
	After *graphql.ID
}) (*RepoBranchesConnectionResolver, error) {
	// fetch repo info
	res, err := loader.RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*ci.RepoInfo)
	if res.(*ci.RepoInfo) == nil {
		return nil, nil
	}

	// calculate the start and end
	start := 0
	if args.After != nil {
		for ; start < len(repo.BranchNames); start++ {
			brID := graphql.ID(r.name + " " + repo.BranchNames[start])
			if brID == *args.After {
				start++
				break
			}
		}
	}
	end := len(repo.BranchNames)
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + int(*args.First)
	}
	if end > len(repo.BranchNames) {
		end = len(repo.BranchNames)
	}

	// build next level resolver
	repoBrRx := &RepoBranchesConnectionResolver{
		pageInfo: &PageInfoResolver{end != len(repo.BranchNames)},
	}
	if start >= end {
		return repoBrRx, nil
	}
	for _, name := range repo.BranchNames[start:end] {
		id := graphql.ID(r.name + " " + name)
		repoBrRx.edges = append(repoBrRx.edges, &RepoBranchesEdgeResolver{
			cursor: id,
			node: &BranchResolver{
				id:       id,
				name:     name,
				repoName: r.name,
			},
		})
	}
	return repoBrRx, nil
}

func (r *RepoResolver) CommitsConnection(args struct {
	First *int32
	After *graphql.ID
}) (*RepoCommitsConnectionResolver, error) {
	// fetch repo info
	res, err := loader.RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*ci.RepoInfo)
	if res.(*ci.RepoInfo) == nil {
		return nil, nil
	}

	// calculate the start and end
	start := 0
	if args.After != nil {
		for ; start < len(repo.CommitHashs); start++ {
			cmID := graphql.ID(r.name + " " + repo.CommitHashs[start])
			if cmID == *args.After {
				start++
				break
			}
		}
	}
	end := len(repo.CommitHashs)
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + int(*args.First)
	}
	if end > len(repo.CommitHashs) {
		end = len(repo.CommitHashs)
	}

	// build next level resolver
	repoCmRx := &RepoCommitsConnectionResolver{
		pageInfo: &PageInfoResolver{end != len(repo.CommitHashs)},
	}
	if start >= end {
		return repoCmRx, nil
	}
	for _, hash := range repo.CommitHashs[start:end] {
		id := graphql.ID(r.name + " " + hash)
		repoCmRx.edges = append(repoCmRx.edges, &RepoCommitsEdgeResolver{
			cursor: id,
			node: &CommitResolver{
				id:       id,
				repoName: r.name,
				hash:     hash,
			},
		})
	}
	return repoCmRx, nil
}

func (r *RepoResolver) RunsConnection(args struct {
	First *int32
	After *graphql.ID
}) (*RepoRunsConnectionResolver, error) {
	// fetch repo info
	res, err := loader.RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*ci.RepoInfo)
	if res.(*ci.RepoInfo) == nil {
		return nil, nil
	}

	// build the num array in descending order
	runNums := make([]int32, 0)
	for n := repo.MaxRunNum; n >= 1; n-- {
		runNums = append(runNums, n)
	}

	// calculate the start and end
	start := 0
	if args.After != nil {
		for ; start < len(runNums); start++ {
			runID := graphql.ID(r.name + " " + strconv.Itoa(int(runNums[start])))
			if runID == *args.After {
				start++
				break
			}
		}
	}
	end := len(runNums)
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + int(*args.First)
	}
	if end > len(runNums) {
		end = len(runNums)
	}

	// build next level resolver
	repoRunRx := &RepoRunsConnectionResolver{
		pageInfo: &PageInfoResolver{end != len(runNums)},
	}
	for _, num := range runNums[start:end] {
		id := graphql.ID(r.name + " " + strconv.Itoa(int(num)))
		repoRunRx.edges = append(repoRunRx.edges, &RepoRunsEdgeResolver{
			cursor: id,
			node: &RunResolver{
				id:       id,
				repoName: r.name,
				num:      num,
			},
		})
	}
	return repoRunRx, nil
}
