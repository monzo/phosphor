package phosphord

import (
	"encoding/json"
	"time"

	log "github.com/cihub/seelog"
	"github.com/golang/protobuf/proto"

	pb "github.com/mondough/phosphor/proto"
)

func (p *PhosphorD) forward(id int) {

	log.Debugf("[Forwarder %v] started", id)

	var (
		b           []byte
		i           int
		decoded     *pb.Annotation
		js          []byte
		buf         = make([][]byte, 0, p.opts.BufferSize)
		metricsTick = time.NewTicker(5 * time.Second)
		timeoutTick = time.NewTicker(time.Duration(p.opts.FlushInterval) * time.Millisecond)
	)

	for {
		select {
		case <-p.exitChan:
			return
		case b = <-p.traceChan:
			i++

			// Log the frame if we're in verbose mode
			if p.opts.Verbose {
				decoded = &pb.Annotation{}
				if err := proto.Unmarshal(b, decoded); err != nil {
					log.Warnf("[Forwarder %v] Couldn't decode trace frame", id)
					continue
				}
				js, _ = json.Marshal(decoded)
				log.Tracef("[Forwarder %v] Received message: %s", id, string(js))
			}

			// Add message to our buffer
			buf = append(buf, b)

			// Forward on if we're at our buffer size
			if len(buf) >= p.opts.BufferSize {
				p.sendTraces(id, &buf)
			}
		case <-timeoutTick.C:
			p.sendTraces(id, &buf)
		case <-metricsTick.C:
			log.Debugf("[Forwarder %v] Processed %v messages", id, i)
		}
	}
}

func (p *PhosphorD) sendTraces(id int, buf *[][]byte) error {
	// Don't publish empty buffers
	if buf == nil || len(*buf) == 0 {
		return nil
	}

	// Attempt to publish
	log.Debugf("[Forwarder %v] Sending %v traces", id, len(*buf))
	if err := p.tr.MultiPublish(*buf); err != nil {
		// we return an error here, but currently ignore it
		// therefore the behaviour will be reattempting to republish the
		// buffer when the next trace arrives to this forwarder
		return err
	}

	// Empty the buffer on success
	*buf = nil

	return nil
}
