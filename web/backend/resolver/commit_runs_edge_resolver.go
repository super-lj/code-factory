package resolver

import graphql "github.com/graph-gophers/graphql-go"

type CommitRunsEdgeResolver struct {
	cursor graphql.ID
	node   *RunResolver
}

func (r *CommitRunsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *CommitRunsEdgeResolver) Node() *RunResolver {
	return r.node
}
