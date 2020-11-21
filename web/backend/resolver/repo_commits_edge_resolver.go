package resolver

import graphql "github.com/graph-gophers/graphql-go"

type RepoCommitsEdgeResolver struct {
	cursor graphql.ID
	node   *CommitResolver
}

func (r *RepoCommitsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

func (r *RepoCommitsEdgeResolver) Node() *CommitResolver {
	return r.node
}
