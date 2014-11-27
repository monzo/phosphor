package forwarder

import (
	"time"

	log "github.com/cihub/seelog"
)

func Start(traceChan chan []byte, numWorkers, bufferSize int) {

	log.Infof("Starting %v forwarders with buffer size of %v", numWorkers, bufferSize)

	for i := 0; i < numWorkers; i++ {
		f := &forwarder{
			id: i,
			ch: traceChan,
		}

		// Do something useful
		go f.work()
	}

}

type forwarder struct {
	id int
	ch chan []byte
}

func (f *forwarder) work() {

	log.Debugf("[Forwarder %v] started", f.id)

	// var b []byte
	var i int
	tick := time.NewTicker(5 * time.Second)

	for {
		// log.Debugf("[Forwarder %v] Waiting for message", f.id)
		select {
		case <-f.ch:
			i++
			// log.Debugf("[Forwarder %v] Received message", f.id)
		case <-tick.C:
			log.Debugf("[Forwarder %v] Processed %v messages", f.id, i)
			i = 0
		}
	}
}
