package store

import (
	"sync"
	"time"

	log "github.com/cihub/seelog"

	"github.com/mondough/phosphor/domain"
)

type MemoryStore struct {
	sync.RWMutex
	traces map[string]*domain.Trace
}

// NewMemoryStore initialises and returns a new MemoryStore
func NewMemoryStore() *MemoryStore {
	s := &MemoryStore{
		traces: make(map[string]*domain.Trace),
	}

	// run stats worker
	go s.statsLoop()

	return s
}

// ReadTrace retrieves a full Trace, composed of Annotations from the store by ID
func (s *MemoryStore) ReadTrace(id string) (*domain.Trace, error) {
	if s == nil {
		return nil, ErrStoreNotInitialised
	}

	s.RLock()
	defer s.RUnlock()

	return s.traces[id], nil
}

// StoreAnnotation into the store, if the trace doesn't not already exist
// this will be created for the global trace ID
func (s *MemoryStore) StoreAnnotation(a *domain.Annotation) error {
	s.Lock()
	defer s.Unlock()

	if s == nil {
		return ErrStoreNotInitialised
	}
	if a == nil {
		return ErrInvalidAnnotation
	}
	if a.TraceId == "" {
		return ErrInvalidTraceId
	}

	// Load our current trace
	t := s.traces[a.TraceId]

	// Initialise a new trace if we don't have it already
	if t == nil {
		t = domain.NewTrace()
	}

	// Add the new annotation to this
	t.AppendAnnotation(a)

	// Store it back
	s.traces[a.TraceId] = t

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
	count := len(s.traces)
	s.RUnlock()

	// Separate processing and logging outside of mutex
	log.Infof("[MemoryStore] Traces stored: %v", count)
}
