package domain

import (
	"errors"
	"sync"
	"time"
)

// NewTrace initialises and returns a new Trace
func NewTrace() *Trace {
	return &Trace{
		Annotation: make([]*Annotation, 0),
	}
}

// Trace represents a full trace of a request
// comprised of a number of Annotations
type Trace struct {
	sync.Mutex

	Annotation []*Annotation `json:"annotations"`
}

// AppendAnnotation to a Trace
func (t *Trace) AppendAnnotation(a *Annotation) error {
	if t == nil {
		return errors.New("Trace is Nil")
	}

	t.Annotation = append(t.Annotation, a)

	return nil
}

// AnnotationType represents an Enum of types of Anotations which Phosphor supports
type AnnotationType int32

const (
	UnknownAnnotationType = AnnotationType(0) // No idea...

	// Calls
	Req     = AnnotationType(1) // Client Request dispatch
	Rsp     = AnnotationType(2) // Client Response received
	In      = AnnotationType(3) // Server Request received
	Out     = AnnotationType(4) // Server Response dispatched
	Timeout = AnnotationType(5) // Client timed out waiting

	// Developer initiated annotations
	// @todo
	// Annotation = AnnotationType(6)
)

// An Annotation represents the smallest individually recorded component of a trace
// These can be assembled into spans, and entire traces of a request to our systems
type Annotation struct {
	TraceId      string // Global Trace Identifier
	SpanId       string // Identifier for this span, non unique - eg. RPC calls would have 4 annotation with this id
	ParentSpanId string // Parent span - eg. nested RPC calls

	Timestamp time.Time     // Timestamp the event occured, can only be compared on the same machine
	Duration  time.Duration // Optional: duration of the event, eg. RPC call

	Hostname    string // Hostname this event originated from
	Origin      string // Fully qualified name of the message origin
	Destination string // Optional: Fully qualified name of the message destination

	AnnotationType AnnotationType // The type of Annotation
	Async          bool           // If the request was fired asynchronously

	Payload     string            // The payload, eg. RPC body, or Annotation
	PayloadSize int32             // Bytes of payload
	KeyValue    map[string]string // Key value debug information
}
