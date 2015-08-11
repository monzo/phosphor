package domain

import (
	"time"

	traceproto "github.com/mondough/phosphor/proto"
)

// AnnotationFromProto converts a proto frame to our domain
func AnnotationFromProto(p *traceproto.Annotation) *Annotation {
	if p == nil {
		return &Annotation{}
	}

	return &Annotation{
		TraceId:        p.TraceId,
		SpanId:         p.SpanId,
		ParentSpanId:   p.ParentId,
		Timestamp:      microsecondInt64ToTime(p.Timestamp),
		Duration:       microsecondInt64ToDuration(p.Duration),
		Hostname:       p.Hostname,
		Origin:         p.Origin,
		Destination:    p.Destination,
		AnnotationType: protoToAnnotationType(p.Type),
		Payload:        p.Payload,
		PayloadSize:    int32(len(p.Payload)),
		KeyValue:       protoToKeyValue(p.KeyValue),
	}
}

// protoToAnnotationType converts a annotation type in our proto to our domain
func protoToAnnotationType(p traceproto.AnnotationType) AnnotationType {
	// Ensure we are within bounds
	at := int32(p)
	if at > 6 || at < 1 {
		at = 0
	}

	return AnnotationType(at)
}

// microsecondInt64ToTime converts an integer number of microseconds
// since the epoch to a time
func microsecondInt64ToTime(i int64) time.Time {
	µsec := i % 1e6
	sec := (i - µsec) / 1e6

	return time.Unix(sec, µsec*1e3)
}

// microsecondInt64ToDuration converts an integer number
// of microseconds to a duration
func microsecondInt64ToDuration(i int64) time.Duration {
	return time.Duration(i) * time.Microsecond
}

// protoToKeyValue converts a repeated set of proto key values
// to a map of keys => values
func protoToKeyValue(p []*traceproto.KeyValue) map[string]string {
	ret := make(map[string]string)
	for _, kv := range p {
		if p == nil {
			continue
		}
		ret[kv.Key] = kv.Value
	}
	return ret
}
