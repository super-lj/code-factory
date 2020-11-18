package main

import (
	code_factory_ci "ci/gen-go/ci"
	"ci/logs"
	"ci/method"
	"ci/util"
	"context"
	"fmt"
)

func IsTargetRepoUpdated(ctx context.Context, r *code_factory_ci.IsTargetRepoUpdatedRequest) (*code_factory_ci.IsTargetRepoUpdatedResponse, error) {
	resp, err := method.IsTargetRepoUpdated(ctx, r)
	logs.CtxInfo(ctx, fmt.Sprintf("IsTargetRepoUpdated request: %s resp: %v", util.MarshallOrElseEmpty(ctx, r), util.MarshallOrElseEmpty(ctx, resp)))
	return resp, err
}

func FetchTargetRepoLastCommit(ctx context.Context, r *code_factory_ci.FetchTargetRepoLastCommitRequest) (*code_factory_ci.FetchTargetRepoLastCommitResonse, error) {
	resp, err := method.FetchTargetRepoLastCommit(ctx, r)
	logs.CtxInfo(ctx, fmt.Sprintf("FetchTargetRepoLastCommit request: %s resp: %v", util.MarshallOrElseEmpty(ctx, r), util.MarshallOrElseEmpty(ctx, resp)))
	return resp, err
}
