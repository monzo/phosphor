package memorystore

import (
	"sync"

	"github.com/mattheath/phosphor/domain"
)

func New() *MemoryStore {
	return &MemoryStore{
		store: make(map[string]Trace),
	}
}

type MemoryStore struct {
	sync.RWMutex
	store map[string]Trace
}

func (s *MemoryStore) GetTrace(id string) (Trace, error) {
	s.RLock()
	defer s.RUnlock()

	return s.store[id], nil
}

func (s *MemoryStore) StoreTraceEvent(e domain.Event) error {
	s.Lock()
	defer s.Unlock()

	// Load our current trace
	t := s.store[e.TraceId]

	// Add the new event to this
	t = append(t, e)

	// Store it back
	s.store[e.TraceId] = t

	return nil
}
