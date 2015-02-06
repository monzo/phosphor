package store

import "github.com/mattheath/phosphor/domain"

type Store interface {
	ReadTrace(id string) (*domain.Trace, error)
	StoreFrame(f *domain.Frame) error
}
