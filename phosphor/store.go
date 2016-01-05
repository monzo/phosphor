package phosphor

import "errors"

type Store interface {
	ReadTrace(id string) (*Trace, error)
	StoreAnnotation(a *Annotation) error
}

var (
	ErrStoreNotInitialised = errors.New("Store is not initialised")
	ErrInvalidAnnotation   = errors.New("Annotation is invalid")
	ErrInvalidTrace        = errors.New("Trace is invalid")
	ErrInvalidTraceId      = errors.New("TraceId is invalid")
)
