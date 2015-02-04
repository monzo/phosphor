package domain

import "time"

// Trace represents a full trace of a request
// comprised of a number of events and annotations
type Trace []Event

// EventType represents an Enum of types of Events which Phosphor can record
type EventType int

const (
	// RPC Calls
	Req = EventType(1) // Client Request dispatch
	Rsp = EventType(2) // Client Response received
	In  = EventType(3) // Server Request received
	Out = EventType(4) // Server Response dispatched

	// Developer initiated annotations
	Annotation = EventType(5)
)

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

	EventType EventType // The type of Event

	Payload     string            // The payload, eg. RPC body, or Annotation
	PayloadSize int32             // Bytes of payload
	KeyValue    map[string]string // Key value debug information
}
