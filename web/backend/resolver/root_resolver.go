package resolver

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/graph-gophers/graphql-go"
)

type RootResolver struct{}

func (r *RootResolver) Repos(ctx context.Context, args struct{ Name *string }) []*RepoResolver {
	var repoRxs []*RepoResolver
	if args.Name != nil {
		repoRxs = append(repoRxs, &RepoResolver{
			id:   graphql.ID(*args.Name),
			name: *args.Name,
		})
	} else {
		repoNameLoader := ctx.Value("loaders").(map[string]*dataloader.Loader)["repo_name"]
		res, err := repoNameLoader.Load(ctx, dataloader.StringKey(""))()
		if err != nil {
			return []*RepoResolver{}
		}
		for _, name := range res.([]string) {
			repoRxs = append(repoRxs, &RepoResolver{
				id:   graphql.ID(name),
				name: name,
			})
		}
	}
	return repoRxs
}
