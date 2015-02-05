package ingester

import (
	"fmt"
	"os"

	nsq "github.com/bitly/go-nsq"
	log "github.com/cihub/seelog"
	"github.com/golang/protobuf/proto"

	"github.com/mattheath/phosphor/domain"
	"github.com/mattheath/phosphor/memorystore"
	traceproto "github.com/mattheath/phosphor/proto"
)

var (
	topic   = "trace"
	channel = "phosphor-server"

	maxInFlight = 200
)

func Run(nsqLookupdHTTPAddrs []string, store *memorystore.MemoryStore) {

	cfg := nsq.NewConfig()
	cfg.UserAgent = fmt.Sprintf("phosphor go-nsq/%s", nsq.VERSION)
	cfg.MaxInFlight = maxInFlight

	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Critical(err)
		os.Exit(1)
	}

	consumer.AddHandler(&IngestionHandler{})

	err = consumer.ConnectToNSQLookupds(nsqLookupdHTTPAddrs)
	if err != nil {
		log.Critical(err)
		os.Exit(1)
	}

	// Block until exit
	<-consumer.StopChan
}

// IngestionHandler exists to match the NSQ handler interface
type IngestionHandler struct{}

// HandleMessage delivered by NSQ
func (ih *IngestionHandler) HandleMessage(message *nsq.Message) error {

	p := &traceproto.TraceFrame{}
	err := proto.Unmarshal(message.Body, p)
	if err != nil {
		// returning an error to NSQ will requeue this
		// failure to unmarshal is permanent
		return nil
	}

	t := domain.FrameFromProto(p)

	log.Debugf("Received trace: %+v", t)

	return nil
}
