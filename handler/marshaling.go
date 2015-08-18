package handler

import (
	"sort"

	"github.com/mondough/phosphor/domain"
	"github.com/mondough/phosphor/proto"
)

func prettyFormatTrace(t *domain.Trace) interface{} {
	return map[string]interface{}{
		"annotations": formatAnnotations(t.Annotation),
	}
}

func formatAnnotations(ans []*domain.Annotation) interface{} {

	sort.Sort(ByTime(ans))

	// Convert to proto
	pa := domain.AnnotationsToProto(ans)

	// Format nicely as JSON
	m := make([]interface{}, 0, len(pa))
	for _, a := range pa {
		m = append(m, formatAnnotation(a))
	}
	return m
}

func formatAnnotation(a *proto.Annotation) interface{} {
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

type ByTime []*domain.Annotation

func (s ByTime) Len() int {
	return len(s)
}
func (s ByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByTime) Less(i, j int) bool {
	return s[i].Timestamp.Before(s[j].Timestamp)
}
