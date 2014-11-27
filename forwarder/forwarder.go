package forwarder

import (
	"time"

	log "github.com/cihub/seelog"
)

func Start(traceChan chan []byte, numWorkers, bufferSize int) {

	log.Infof("Starting %v forwarders with buffer size of %v", numWorkers, bufferSize)

	for i := 0; i < numWorkers; i++ {
		f := &forwarder{
			id:            i,
			ch:            traceChan,
			messageBuffer: make([][]byte, 0, bufferSize),
			bufferSize:    bufferSize,
		}

		// Do something useful
		go f.work()
	}

}

type forwarder struct {
	id            int
	ch            chan []byte
	messageBuffer [][]byte
	bufferSize    int
}

func (f *forwarder) work() {

	log.Debugf("[Forwarder %v] started", f.id)

	var b []byte
	var i int

	metricsTick := time.NewTicker(5 * time.Second)
	timeoutTick := time.NewTicker(2 * time.Second)

	for {
		// log.Debugf("[Forwarder %v] Waiting for message", f.id)
		select {
		case b = <-f.ch:
			i++
			// log.Debugf("[Forwarder %v] Received message", f.id)

			// Add message to our buffer
			f.messageBuffer = append(f.messageBuffer, b)

			// Forward on if we're at our buffer size
			if len(f.messageBuffer) >= f.bufferSize {
				f.send()
			}
		case <-timeoutTick.C:
			f.send()
		case <-metricsTick.C:
			log.Debugf("[Forwarder %v] Processed %v messages", f.id, i)
			// i = 0
		}
	}
}

func (f *forwarder) send() {

	log.Infof("[Forwarder %v] Sent %v messages", f.id, len(f.messageBuffer))

	// Empty the buffer
	f.messageBuffer = nil
}
