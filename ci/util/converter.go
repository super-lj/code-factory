package util

import (
	"ci/domain"
	code_factory_ci "ci/thrift_gen/code/factory/ci"
)

func ConvertCommit(commit *domain.Commit) *code_factory_ci.CommitStruct {
	resp := &code_factory_ci.CommitStruct{
		Id:             &(commit.Id),
		Msg:            &(commit.Msg),
		Author:         &(commit.Author),
		LastUpdateTime: &(commit.LastUpdateTime),
	}
	return resp
}