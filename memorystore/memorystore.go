package memorystore

import (
	"sync"

	"github.com/mattheath/phosphor/domain"
)

func New() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]domain.Trace),
	}
}

type MemoryStore struct {
	sync.RWMutex
	store map[string]domain.Trace
}

func (s *MemoryStore) GetTrace(id string) (domain.Trace, error) {
	s.RLock()
	defer s.RUnlock()

	return s.store[id], nil
}

func (s *MemoryStore) StoreTraceFrame(f domain.Frame) error {
	s.Lock()
	defer s.Unlock()

	// Load our current trace
	t := s.store[e.TraceId]

	// Add the new frame to this
	t = append(t, f)

	// Store it back
	s.store[e.TraceId] = t

	return nil
}
