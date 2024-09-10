package trace_ctx

import (
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	SpanID string

	Span struct {
		traceID      TraceID
		id           SpanID
		parentSpanID *SpanID
		operation    string
		startTime    int64
		duration     int64
		tags         *map[string]string
		attrs        []slog.Attr
		mtx          *sync.Mutex
	}
)

func NewSpan(traceID TraceID, operation string, parentID *SpanID) *Span {
	return &Span{
		id:           genSpanID(),
		traceID:      traceID,
		parentSpanID: parentID,
		operation:    operation,
		startTime:    time.Now().UnixNano() / int64(time.Microsecond),
		tags:         new(map[string]string),
		attrs:        []slog.Attr{},
		mtx:          &sync.Mutex{},
	}
}

func (s *Span) End() {
	s.duration = (time.Now().UnixNano() / int64(time.Microsecond)) - s.startTime
}

func (s *Span) AttrAdd(a slog.Attr) *Span {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.attrs = append(s.attrs, a)
	return s
}

func (s *Span) GetAttrs() (a []slog.Attr) {
	a = append(a, s.attrs...)
	a = append(a,
		slog.String(TraceIDKeyName, string(s.traceID)),
		slog.String("span_id", string(s.id)),
		slog.String("operation_name", s.operation),
		slog.Int64("start_time", s.startTime),
		slog.Int64("duration", s.duration),
	)
	if s.parentSpanID != nil {
		a = append(a, slog.String("parent_span_id", string(*s.parentSpanID)))
	}
	if s.tags != nil {
		a = append(a, slog.Any("tags", s.tags))
	}
	return a
}

func genSpanID() SpanID {
	return SpanID(replace(uuid.New().String()))
}
