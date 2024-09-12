package trace_ctx

import (
	"sync"

	"github.com/google/uuid"
)

type (
	TraceID string

	Trace struct {
		ID    TraceID
		spans []*Span
		mtx   *sync.Mutex
	}
)

func NewTraceWithID(id TraceID) *Trace {
	if len(id) == 0 {
		panic("id is empty")
	}
	return &Trace{
		ID:    id,
		spans: []*Span{},
		mtx:   &sync.Mutex{},
	}
}

func NewTrace() *Trace {
	return &Trace{
		ID:    genTraceID(),
		spans: []*Span{},
		mtx:   &sync.Mutex{},
	}
}

func (t *Trace) StartSpan(operation string, parent *Span) *Span {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	parentID := new(SpanID)
	if parent != nil {
		parentID = &parent.id
	}

	span := NewSpan(t.ID, operation, parentID)
	t.spans = append(t.spans, span)
	return span
}

func (t *Trace) WithSpan(operation string, parent *Span, f func()) *Span {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	parentID := new(SpanID)
	if parent != nil {
		parentID = &parent.id
	}

	span := NewSpan(t.ID, operation, parentID)
	f()
	span.End()
	t.spans = append(t.spans, span)

	return span
}

func genTraceID() TraceID {
	return TraceID(replace(uuid.NewString()))
}
