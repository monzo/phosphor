package domain

import (
	"time"

	traceproto "github.com/mattheath/phosphor/proto"
)

// FrameFromProto converts a proto frame to our domain
func FrameFromProto(p *traceproto.TraceFrame) *Frame {
	return &Frame{
		TraceId:      p.GetTraceId(),
		SpanId:       p.GetSpanId(),
		ParentSpanId: p.GetParentId(),
		Timestamp:    nanosecondInt64ToTime(p.GetTimestamp()),
		Duration:     nanosecondInt64ToDuration(p.GetDuration()),
		Hostname:     p.GetHostname(),
		Origin:       p.GetOrigin(),
		Destination:  p.GetDestination(),
		FrameType:    protoToFrameType(p.GetType()),
		Payload:      p.GetPayload(),
		PayloadSize:  int32(len(p.GetPayload())),
		KeyValue:     protoToKeyValue(p.GetKeyValue()),
	}
}

// protoToFrameType converts a frametype in our proto to our domain
func protoToFrameType(p traceproto.FrameType) FrameType {
	// Ensure we are within bounds
	ft := int32(p)
	if ft > 6 || ft < 1 {
		ft = 0
	}

	return FrameType(ft)
}

// nanosecondInt64ToTime converts an integer number of nanoseconds
// since the epoch to a time
func nanosecondInt64ToTime(i int64) time.Time {
	nsec := i % 1e9
	sec := (i - nsec) / 1e9

	return time.Unix(sec, nsec)
}

// nanosecondInt64ToDuration converts an integer number
// of nanoseconds to a duration
func nanosecondInt64ToDuration(i int64) time.Duration {
	return time.Duration(i) * time.Nanosecond
}

// protoToKeyValue converts a repeated set of proto key values
// to a map of keys => values
func protoToKeyValue(p []*traceproto.KeyValue) map[string]string {
	ret := make(map[string]string)
	for _, kv := range p {
		ret[kv.GetKey()] = kv.GetValue()
	}
	return ret
}
