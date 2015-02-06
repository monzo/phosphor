package store

import (
	"sync"
	"time"

	log "github.com/cihub/seelog"

	"github.com/mattheath/phosphor/domain"
)

type MemoryStore struct {
	sync.RWMutex
	store map[string]*domain.Trace
}

// NewMemoryStore initialises and returns a new MemoryStore
func NewMemoryStore() *MemoryStore {
	s := &MemoryStore{
		store: make(map[string]*domain.Trace),
	}

	// run stats worker
	go s.statsLoop()

	return s
}

// ReadTrace retrieves a full Trace, composed of Frames from the store by ID
func (s *MemoryStore) ReadTrace(id string) (*domain.Trace, error) {
	if s == nil {
		return nil, ErrStoreNotInitialised
	}

	s.RLock()
	defer s.RUnlock()

	return s.store[id], nil
}

// StoreTraceFrame into the store, if the trace doesn't not already exist
// this will be created for the global trace ID
func (s *MemoryStore) StoreFrame(f *domain.Frame) error {
	s.Lock()
	defer s.Unlock()

	if s == nil {
		return ErrStoreNotInitialised
	}
	if f == nil {
		return ErrNilFrame
	}
	if f.TraceId == "" {
		return ErrInvalidTraceId
	}

	// Load our current trace
	t := s.store[f.TraceId]

	// Initialise a new trace if we don't have it already
	if t == nil {
		t = domain.NewTrace()
	}

	// Add the new frame to this
	t.AppendFrame(f)

	// Store it back
	s.store[f.TraceId] = t

	return nil
}

// statsLoop loops and outputs stats every 5 seconds
func (s *MemoryStore) statsLoop() {

	tick := time.NewTicker(5 * time.Second)

	// @todo listen for shutdown, stop ticker and exit cleanly
	for {
		<-tick.C // block until tick

		s.printStats()
	}
}

// printStats about the status of the memorystore to stdout
func (s *MemoryStore) printStats() {

	// Get some data while under the mutex
	s.RLock()
	count := len(s.store)
	s.RUnlock()

	// Separate processing and logging outside of mutex
	log.Infof("[MemoryStore] Traces stored: %v", count)
}
