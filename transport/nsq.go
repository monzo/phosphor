package transport

import (
	"errors"
	"math/rand"

	nsq "github.com/bitly/go-nsq"
	log "github.com/cihub/seelog"

	"github.com/mattheath/phosphor/util"
)

var (
	ErrPublishFailure    = errors.New("Failed to publish to NSQD")
	ErrNoConfiguredNodes = errors.New("No NSQD nodes are configured")
)

// NewNSQTransport initialises a Transport over NSQ
func NewNSQTransport(topic string, nsqdTCPAddrs util.StringArray) (Transport, error) {

	// Currently using default config
	cfg := nsq.NewConfig()

	// Create a producer for each nsqd node provided
	producers := make(map[string]*nsq.Producer)
	producersIndex := make([]*nsq.Producer, 0, len(nsqdTCPAddrs))
	for _, addr := range nsqdTCPAddrs {
		producer, err := nsq.NewProducer(addr, cfg)
		if err != nil {
			log.Warnf("failed to create nsq.Producer - %s", err)
		}
		producers[addr] = producer
		producersIndex = append(producersIndex, producers[addr])
	}

	return &NSQPublisher{
		topic:          topic,
		producers:      producers,
		producersIndex: producersIndex,
	}, nil
}

type NSQPublisher struct {
	topic          string
	producers      map[string]*nsq.Producer
	producersIndex []*nsq.Producer
}

func (p *NSQPublisher) MultiPublish(body [][]byte) error {

	if len(p.producers) == 0 {
		return ErrNoConfiguredNodes
	}

	// Round robin, from a random starting position
	i := rand.Intn(len(p.producers)) - 1

	// Attempt up to our number of configured nodes
	for attempt := 0; attempt < len(p.producers); attempt++ {

		// Move to next host, or cycle back around
		i++
		if i >= len(p.producers) {
			i = 0
		}

		// Attempt to publish
		pd := p.producersIndex[i]
		if err := pd.MultiPublish(p.topic, body); err == nil {
			// success!
			return nil
		}
	}

	// We've run out of nodes, and not managed to publish
	return ErrPublishFailure
}
