package resolver

type PageInfoResolver struct {
	hasNextPage bool
}

func (r *PageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}
