package store

import (
	"errors"

	"github.com/mattheath/phosphor/domain"
)

type Store interface {
	ReadTrace(id string) (*domain.Trace, error)
	StoreAnnotation(a *domain.Annotation) error
}

var (
	ErrStoreNotInitialised = errors.New("Store is not initialised")
	ErrInvalidAnnotation   = errors.New("Annotation is invalid")
	ErrInvalidTrace        = errors.New("Trace is invalid")
	ErrInvalidTraceId      = errors.New("TraceId is invalid")
)
