package transport

import (
	nsq "github.com/bitly/go-nsq"
	log "github.com/cihub/seelog"

	"github.com/mattheath/phosphor/util"
)

// NewNSQTransport initialises a Transport over NSQ
func NewNSQTransport(topic string, nsqdTCPAddrs util.StringArray) (Transport, error) {

	// Currently using default config
	cfg := nsq.NewConfig()

	// Create a producer for each nsqd node provided
	producers := make(map[string]*nsq.Producer)
	for _, addr := range nsqdTCPAddrs {
		producer, err := nsq.NewProducer(addr, cfg)
		if err != nil {
			log.Warnf("failed to create nsq.Producer - %s", err)
		}
		producers[addr] = producer
	}

	return &NSQPublisher{
		topic:     topic,
		producers: producers,
	}, nil
}

type NSQPublisher struct {
	topic     string
	producers map[string]*nsq.Producer
}

func (p *NSQPublisher) MultiPublish(body [][]byte) error {
	return nil
}
