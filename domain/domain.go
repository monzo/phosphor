package domain

import "time"

// Trace represents a full trace of a request
// comprised of a number of frames
type Trace []Frame

// FrameType represents an Enum of types of Events which Phosphor can record
type FrameType int

const (
	// Calls
	Req     = FrameType(1) // Client Request dispatch
	Rsp     = FrameType(2) // Client Response received
	In      = FrameType(3) // Server Request received
	Out     = FrameType(4) // Server Response dispatched
	Timeout = FrameType(5) // Client timed out waiting

	// Developer initiated annotations
	Annotation = FrameType(6)
)

// A Frame represents the smallest individually fired component of a trace
// These can be assembled into spans, and entire traces of a request to our systems
type Frame struct {
	TraceId      string // Global Trace Identifier
	SpanId       string // Identifier for this span, non unique - eg. RPC calls would have 4 frames with this id
	ParentSpanId string // Parent span - eg. nested RPC calls

	Timestamp time.Time     // Timestamp the event occured, can only be compared on the same machine
	Duration  time.Duration // Optional: duration of the event, eg. RPC call

	Hostname    string // Hostname this event originated from
	Origin      string // Fully qualified name of the message origin
	Destination string // Optional: Fully qualified name of the message destination

	EventType EventType // The type of Event

	Payload     string            // The payload, eg. RPC body, or Annotation
	PayloadSize int32             // Bytes of payload
	KeyValue    map[string]string // Key value debug information
}
