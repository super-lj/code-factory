package resolver

type CommitRunsConnectionResolver struct {
	edges    []*CommitRunsEdgeResolver
	pageInfo *PageInfoResolver
}

func (r *CommitRunsConnectionResolver) Edges() []*CommitRunsEdgeResolver {
	return r.edges
}

func (r *CommitRunsConnectionResolver) PageInfo() *PageInfoResolver {
	return r.pageInfo
}
