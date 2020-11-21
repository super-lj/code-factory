package resolver

type RepoBranchesConnectionResolver struct {
	edges    []*RepoBranchesEdgeResolver
	pageInfo *PageInfoResolver
}

func (r *RepoBranchesConnectionResolver) Edges() []*RepoBranchesEdgeResolver {
	return r.edges
}

func (r *RepoBranchesConnectionResolver) PageInfo() *PageInfoResolver {
	return r.pageInfo
}
