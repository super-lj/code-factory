package resolver

import graphql "github.com/graph-gophers/graphql-go"

type BranchRunsEdgeResolver struct {
	cursor graphql.ID
	node   *RunResolver
}

func (r *BranchRunsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *BranchRunsEdgeResolver) Node() *RunResolver {
	return r.node
}
