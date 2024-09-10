package trace_ctx_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/rzaripov1990/trace_ctx"
)

func TestTracing(t *testing.T) {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		}),
	)

	trace := trace_ctx.NewTrace()
	ctx := trace_ctx.WithTrace(context.Background(), trace)

	span1 := trace.StartSpan("operation-1", nil)
	time.Sleep(100 * time.Millisecond)
	span1.End()
	span1.AttrAdd(slog.Bool("slow", true))
	log.LogAttrs(ctx, slog.LevelDebug, "operation-1", span1.GetAttrs()...)

	span2 := trace.StartSpan("operation-2", span1)
	time.Sleep(200 * time.Millisecond)
	span2.End()
	log.LogAttrs(ctx, slog.LevelDebug, "operation-2", span2.GetAttrs()...)

	var (
		err   error
		count int
	)
	log.LogAttrs(
		ctx,
		slog.LevelDebug,
		"operation-3",
		trace.WithSpan("operation-3", span2,
			func() {
				if count == 0 && err == nil {
					count = 5
				}
				time.Sleep(300 * time.Millisecond)
			},
		).AttrAdd(slog.Int("count", count)).GetAttrs()...)

	fmt.Println("Spans logged to standard output")

	//fmt.Printf("%#v", trace_ctx.GetTrace(ctx))
}
