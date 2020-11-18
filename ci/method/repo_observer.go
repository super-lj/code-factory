package method

import (
	"ci/dao"
	"ci/domain"
	"ci/logs"
	code_factory_ci "ci/gen-go/ci"
	"ci/util"
	"context"
	"fmt"
	git "github.com/libgit2/git2go"
	"time"
)

func IsTargetRepoUpdated(ctx context.Context, r *code_factory_ci.IsTargetRepoUpdatedRequest) (*code_factory_ci.IsTargetRepoUpdatedResponse, error) {
	repoStruct, opRes := dao.GetRepoInfoById(ctx, r.GetHash())
	if !opRes.Success() {
		logs.CtxError(ctx, "[GetRepoInfoById] fail")
		return buildIsTargetRepoUpdated(false, "[GetRepoInfoById] fail", util.ErrDBRead), nil
	}
	repo, _ := git.OpenRepository(repoStruct.Url)
	remote, _ := repo.Remotes.Lookup("origin")
	if err := remote.Fetch([]string{}, nil, ""); err != nil {
		logs.CtxError(ctx, fmt.Sprintf("fetch remote fail, err: %+v", err))
		return buildIsTargetRepoUpdated(false, fmt.Sprintf("fetch remote fail, err: %+v", err), util.ErrSystemInternal), nil
	}
	head, err := repo.Head()
	if err != nil {
		logs.CtxError(ctx, fmt.Sprintf("get remote head fail, err: %+v", err))
		return buildIsTargetRepoUpdated(false, fmt.Sprintf("get remote head fail, err: %+v", err), util.ErrSystemInternal), nil
	}
	commit, err := repo.LookupCommit(head.Branch().Target())
	if err != nil {
		logs.CtxError(ctx, fmt.Sprintf("get latest commit fail, err: %+v", err))
		return buildIsTargetRepoUpdated(false, fmt.Sprintf("get latest commit fail, err: %+v", err), util.ErrSystemInternal), nil
	}
	c := &domain.Commit{
		Hash:           commit.Id().String(),
		Msg:            commit.Message(),
		Author:         commit.Author().Name,
		RunNums: 		time.Now().UnixNano() % 1e6 / 1e3,
	}
	dao.SaveCommitToCache(c)
	return buildSucIsTargetRepoUpdated(), nil
}

func FetchTargetRepoLastCommit(ctx context.Context, r *code_factory_ci.FetchTargetRepoLastCommitRequest) (*code_factory_ci.FetchTargetRepoLastCommitResonse, error) {
	commit, opRes := dao.GetCommitByHashFromCache(r.GetHash())
	if !opRes.Success() { // ignore cache miss
		logs.CtxWarn(ctx, "[GetCommitByIdFromCache] fail")
		return buildFetchTargetRepoLastCommitResp(nil, "[GetCommitByIdFromCache] fail", util.ErrRedis), nil
	}
	return buildSucFetchTargetRepoLastCommitResp(util.ConvertCommit(commit)), nil
}

func buildSucIsTargetRepoUpdated() *code_factory_ci.IsTargetRepoUpdatedResponse {
	resp := buildIsTargetRepoUpdated(true, "", util.Success)
	return resp
}

func buildIsTargetRepoUpdated(isUpdated bool, msg string, code int32) *code_factory_ci.IsTargetRepoUpdatedResponse {
	resp := &code_factory_ci.IsTargetRepoUpdatedResponse{
		IsUpdated: &isUpdated,
		BaseResp: &code_factory_ci.BaseResp{
			StatusMessage: msg,
			StatusCode:    code,
			Extra:         nil,
		},
	}
	return resp
}

func buildSucFetchTargetRepoLastCommitResp(commit *code_factory_ci.CommitInfo) *code_factory_ci.FetchTargetRepoLastCommitResonse {
	resp := buildFetchTargetRepoLastCommitResp(commit, "", util.Success)
	return resp
}

func buildFetchTargetRepoLastCommitResp(commit *code_factory_ci.CommitInfo, msg string, code int32) *code_factory_ci.FetchTargetRepoLastCommitResonse {
	resp := &code_factory_ci.FetchTargetRepoLastCommitResonse{
		Commit: commit,
		BaseResp: &code_factory_ci.BaseResp{
			StatusMessage: msg,
			StatusCode:    code,
			Extra:         nil,
		},
	}
	return resp
}
