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

func TestCtx1(t *testing.T) {
	mainCtx := context.Background()
	trace := trace_ctx.GetTrace(mainCtx)
	ctx := trace_ctx.WithTrace(mainCtx, trace)

	fmt.Println(trace_ctx.GetTraceID(ctx))
}
