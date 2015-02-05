package memorystore

import (
	"sync"
	"time"

	"github.com/mattheath/phosphor/domain"
)

func New() *MemoryStore {
	s := &MemoryStore{
		store: make(map[string]domain.Trace),
	}

	// run stats worker
	go s.statsLoop()

	return s
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

func (s *MemoryStore) statsLoop() {

	tick := time.NewTicker(5 * time.Second)

	// @todo listen for shutdown, stop ticker and exit cleanly
	for {
		<-tick.C // block until tick

		s.printStats()
	}
}

func (s *MemoryStore) printStats() {

	// Get some data while under the mutex
	s.RLock()
	count := len(s.store)
	s.RUnlock()

	// Separate processing and logging outside of mutex
	log.Infof("[Phosphor] Traces stored: %v", count)
}
