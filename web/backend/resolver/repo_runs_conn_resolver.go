package resolver

type RepoRunsConnectionResolver struct {
	edges    []*RepoRunsEdgeResolver
	pageInfo *PageInfoResolver
}

func (r *RepoRunsConnectionResolver) Edges() []*RepoRunsEdgeResolver {
	return r.edges
}

func (r *RepoRunsConnectionResolver) PageInfo() *PageInfoResolver {
	return r.pageInfo
}
