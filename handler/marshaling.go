package handler

import (
	"github.com/mondough/phosphor/domain"
	"github.com/mondough/phosphor/proto"
)

func prettyFormatTrace(t *domain.Trace) interface{} {
	return map[string]interface{}{
		"annotations": formatAnnotations(t.Annotation),
	}
}

func formatAnnotations(ans []*domain.Annotation) interface{} {
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
		"timestamp":   a.Timestamp,
		"duration":    a.Duration,
		"hostname":    a.Hostname,
		"origin":      a.Origin,
		"destination": a.Destination,
		"payload":     a.Payload,
		"key_value":   a.KeyValue,
	}
}
