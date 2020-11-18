package util

import (
	"ci/domain"
	code_factory_ci "ci/gen-go/ci"
)

func ConvertCommit(commit *domain.Commit) *code_factory_ci.CommitInfo {
	resp := &code_factory_ci.CommitInfo{
		Hash:    &(commit.Hash),
		Msg:     &(commit.Msg),
		Author:  &(commit.Author),
		RunNums: commit.RunNums,
	}
	return resp
}
