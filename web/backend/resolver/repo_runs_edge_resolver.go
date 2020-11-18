package resolver

import graphql "github.com/graph-gophers/graphql-go"

type RepoRunsEdgeResolver struct {
	cursor graphql.ID
	node   *RunResolver
}

func (r *RepoRunsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *RepoRunsEdgeResolver) Node() *RunResolver {
	return r.node
}
