package resolver

import "github.com/graph-gophers/graphql-go"

type RepoBranchesEdgeResolver struct {
	cursor graphql.ID
	node   *BranchResolver
}

func (r *RepoBranchesEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *RepoBranchesEdgeResolver) Node() *BranchResolver {
	return r.node
}
