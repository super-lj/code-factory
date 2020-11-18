package resolver

type BranchRunsConnectionResolver struct {
	edges    []*BranchRunsEdgeResolver
	pageInfo *PageInfoResolver
}

func (r *BranchRunsConnectionResolver) Edges() []*BranchRunsEdgeResolver {
	return r.edges
}

func (r *BranchRunsConnectionResolver) PageInfo() *PageInfoResolver {
	return r.pageInfo
}
