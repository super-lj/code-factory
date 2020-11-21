package resolver

type RepoCommitsConnectionResolver struct {
	edges    []*RepoCommitsEdgeResolver
	pageInfo *PageInfoResolver
}

func (r *RepoCommitsConnectionResolver) Edges() []*RepoCommitsEdgeResolver {
	return r.edges
}

func (r *RepoCommitsConnectionResolver) PageInfo() *PageInfoResolver {
	return r.pageInfo
}
