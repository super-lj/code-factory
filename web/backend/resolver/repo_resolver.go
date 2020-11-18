package resolver

import (
	"context"
	"strconv"
	"strings"
	"web-backend/mock"

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
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*mock.RepoInfo)
	if res.(*mock.RepoInfo) == nil {
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
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*mock.RepoInfo)
	if res.(*mock.RepoInfo) == nil {
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
	res, err := RepoInfoloader.Load(context.TODO(), dataloader.StringKey(r.name))()
	if err != nil {
		return nil, err
	}
	repo := res.(*mock.RepoInfo)
	if res.(*mock.RepoInfo) == nil {
		return nil, nil
	}

	// calculate the start and end
	var start int32 = 1
	if args.After != nil {
		strs := strings.Split(string(*args.After), " ")
		if len(strs) != 2 {
			return nil, nil
		}
		afterNumInt64, err := strconv.ParseInt(strs[1], 10, 32)
		if err != nil {
			return nil, err
		}
		afterNum := int32(afterNumInt64)
		if afterNum >= repo.MaxRunNum {
			repoRunRx := &RepoRunsConnectionResolver{
				pageInfo: &PageInfoResolver{true},
			}
			return repoRunRx, nil
		}
		start = afterNum + 1
	}
	end := repo.MaxRunNum + 1
	if args.First != nil {
		if *args.First < 0 {
			return nil, nil
		}
		end = start + *args.First
	}
	if end > repo.MaxRunNum+1 {
		end = repo.MaxRunNum + 1
	}

	// build next level resolver
	repoRunRx := &RepoRunsConnectionResolver{
		pageInfo: &PageInfoResolver{end != repo.MaxRunNum+1},
	}
	for num := start; num < end; num++ {
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
