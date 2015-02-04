package domain

import "time"

// Trace represents a full trace of a request
// comprised of a number of events and annotations
type Trace struct {
	Events      []Event
	Annotations []Annotation
}

// An Event represents a section of an RPC call between systems
type Event struct {
	TraceId       string // Global Trace Identifier
	EventId       string // Identifier for this event, non unique - eg. RPC calls would have 4 of these
	ParentEventId string // Parent event - eg. nested RPC calls

	Timestamp time.Time     // Timestamp the event occured, can only be compared on the same machine
	Duration  time.Duration // Optional: duration of the event, eg. RPC call

	Hostname    string // Hostname this event originated from
	Origin      string // Fully qualified name of the message origin
	Destination string // Fully qualified name of the message destination

	EventType   string // The type of Event
	Payload     string // The returned payload (if possible)
	PayloadSize int32  // Bytes of payload
}

// Annotation describes an optional additional annotation which occurs within a call
type Annotation struct {
	TraceId       string
	EventId       string
	ParentEventId string

	Timestamp time.Time // Timestamp the event occured, can only be compared on the same machine

	Hostname string // Hostname this event originated from

	Content  string            // Arbitrary content, eg. debug output
	KeyValue map[string]string // Key value debug information
}
