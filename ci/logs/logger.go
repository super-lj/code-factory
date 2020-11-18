package logs

import (
	"context"
	"fmt"
)

func CtxInfo(ctx context.Context, v interface{}) string {
	return fmt.Sprintf("[Info] %v, %v", ctx, v)
}

func CtxWarn(ctx context.Context, v interface{}) string {
	return fmt.Sprintf("[Warn] %v, %v", ctx, v)
}

func CtxError(ctx context.Context, v interface{}) string {
	return fmt.Sprintf("[Error] %v, %v", ctx, v)
}
