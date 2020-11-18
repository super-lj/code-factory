package util

import (
	"ci/logs"
	"context"
	"encoding/json"
	"fmt"
)

func MarshallOrElseEmpty(ctx context.Context, v interface{}) string {
	if v == nil {
		return "nil"
	}
	data, err := json.Marshal(v)
	if err != nil {
		logs.CtxError(ctx, fmt.Sprintf("marshall exception, err: %v", err))
	}
	return string(data)
}
