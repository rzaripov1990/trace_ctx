// Package `trace_ctx` enables precise control over the
// setting and retrieval of the `trace_id` within a context,
// ensuring that the unique identifier is consistently managed
// throughout the lifecycle of a request.
//
// This capability is essential for maintaining traceability
// and correlating log entries across different components
// of the application.

package trace_ctx

import (
	"context"
)

var (
	TraceKeyInCtx any = new(TraceID)
	TraceInCtx    any = new(Trace)
)

const (
	TraceIDKeyName = "trace_id"
)

func GetTraceID(ctx context.Context) TraceID {
	if ctx == nil {
		return genTraceID()
	}

	val := ctx.Value(TraceKeyInCtx)
	if val != nil {
		return val.(TraceID)
	}

	return genTraceID()
}

func WithTraceID(ctx context.Context) context.Context {
	if ctx == nil {
		panic("ctx is nil")
	}

	return context.WithValue(ctx, TraceKeyInCtx, genTraceID())
}

func SetTraceID(ctx context.Context, traceID TraceID) context.Context {
	if ctx == nil {
		panic("ctx is nil")
	}

	return context.WithValue(ctx, TraceKeyInCtx, traceID)
}

func WithTrace(ctx context.Context, trace *Trace) context.Context {
	if ctx == nil {
		panic("ctx is nil")
	}

	return context.WithValue(ctx, TraceInCtx, trace)
}

func GetTrace(ctx context.Context) *Trace {
	if ctx == nil {
		panic("ctx is nil")
	}

	val := ctx.Value(TraceInCtx)
	if val != nil {
		return val.(*Trace)
	}

	return nil
}
