package store

import (
	"errors"

	"github.com/mattheath/phosphor/domain"
)

type Store interface {
	ReadTrace(id string) (*domain.Trace, error)
	StoreFrame(f *domain.Frame) error
}

var (
	ErrStoreNotInitialised = errors.New("Store is not initialised")
	ErrInvalidFrame        = errors.New("Frame is invalid")
	ErrInvalidTrace        = errors.New("Trace is invalid")
	ErrInvalidTraceId      = errors.New("TraceId is invalid")
)
