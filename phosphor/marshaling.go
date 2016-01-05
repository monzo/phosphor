package phosphor

import (
	"sort"
	"time"

	"github.com/mondough/phosphor/proto"
)

func prettyFormatTrace(t *Trace) interface{} {
	return map[string]interface{}{
		"annotations": formatAnnotations(t.Annotation),
	}
}

func formatAnnotations(ans []*Annotation) interface{} {

	sort.Sort(ByTime(ans))

	// Convert to proto
	pa := AnnotationsToProto(ans)

	// Format nicely as JSON
	m := make([]interface{}, 0, len(pa))
	for _, a := range pa {
		m = append(m, formatAnnotation(a))
	}
	return m
}

func formatAnnotation(a *traceproto.Annotation) interface{} {
	return map[string]interface{}{
		"trace_id":    a.TraceId,
		"span_id":     a.SpanId,
		"parent_id":   a.ParentId,
		"type":        a.Type.String(),
		"async":       a.Async,
		"timestamp":   a.Timestamp,
		"duration":    a.Duration,
		"hostname":    a.Hostname,
		"origin":      a.Origin,
		"destination": a.Destination,
		"payload":     a.Payload,
		"key_value":   a.KeyValue,
	}
}

type ByTime []*Annotation

func (s ByTime) Len() int {
	return len(s)
}
func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByTime) Less(i, j int) bool {
	return s[i].Timestamp.Before(s[j].Timestamp)
}

// ---

// ProtoToAnnotation converts a proto annotation to our domain
func ProtoToAnnotation(p *traceproto.Annotation) *Annotation {
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
		Async:          p.Async,
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

// annotationTypeToProto converts a annotation type in our domain to proto format
func annotationTypeToProto(at AnnotationType) traceproto.AnnotationType {
	// Ensure we are within bounds
	p := int32(at)
	if p > 6 || p < 1 {
		p = 0
	}

	return traceproto.AnnotationType(p)
}

// microsecondInt64ToTime converts an integer number of microseconds
// since the epoch to a time
func microsecondInt64ToTime(i int64) time.Time {
	µsec := i % 1e6
	sec := (i - µsec) / 1e6

	return time.Unix(sec, µsec*1e3)
}

// timeToMicrosecondInt64 converts a time to µseconds since epoch as int64
func timeToMicrosecondInt64(t time.Time) int64 {
	sec := t.Unix() * 1e6
	µsec := int64(t.Nanosecond() / 1e3)

	return sec + µsec
}

// microsecondInt64ToDuration converts an integer number
// of microseconds to a duration
func microsecondInt64ToDuration(i int64) time.Duration {
	return time.Duration(i) * time.Microsecond
}

// durationToMicrosecondInt64 returns a duration to the nearest µs
func durationToMicrosecondInt64(d time.Duration) int64 {
	return d.Nanoseconds() / 1e3
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

// keyValueToProto converts a map of keys => values to a  repeated set
// of proto key values
func keyValueToProto(m map[string]string) []*traceproto.KeyValue {
	ret := make([]*traceproto.KeyValue, 0, len(m))
	for k, v := range m {
		kv := &traceproto.KeyValue{
			Key:   k,
			Value: v,
		}
		ret = append(ret, kv)
	}
	return ret
}

// AnnotationsToProto converts a slice of domain annotations to our proto format
func AnnotationsToProto(a []*Annotation) []*traceproto.Annotation {
	ret := make([]*traceproto.Annotation, 0, len(a))
	for _, v := range a {
		ret = append(ret, AnnotationToProto(v))
	}
	return ret
}

// AnnotationToProto converts a domain annotation to our proto format
func AnnotationToProto(a *Annotation) *traceproto.Annotation {
	if a == nil {
		return &traceproto.Annotation{}
	}

	return &traceproto.Annotation{
		TraceId:  a.TraceId,
		SpanId:   a.SpanId,
		ParentId: a.ParentSpanId,
		Type:     annotationTypeToProto(a.AnnotationType),
		Async:    a.Async,

		Timestamp: timeToMicrosecondInt64(a.Timestamp),
		Duration:  durationToMicrosecondInt64(a.Duration),

		Hostname:    a.Hostname,
		Origin:      a.Origin,
		Destination: a.Destination,
		Payload:     a.Payload,
		KeyValue:    keyValueToProto(a.KeyValue),
	}
}
