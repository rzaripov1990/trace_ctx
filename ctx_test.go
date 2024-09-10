package trace_ctx_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/rzaripov1990/trace_ctx"
)

func TestCtx(t *testing.T) {
	ctx := trace_ctx.WithTraceID(context.Background())
	fmt.Println(trace_ctx.GetTraceID(ctx))
}
