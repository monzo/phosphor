package phosphor

import (
	"fmt"
	"os"

	nsq "github.com/bitly/go-nsq"
	log "github.com/cihub/seelog"
	"github.com/golang/protobuf/proto"

	traceproto "github.com/mondough/phosphor/proto"
)

var (
	topic   = "trace"
	channel = "phosphor-server"

	maxInFlight = 200
	concurrency = 10
)

// Run the trace ingester, ingesting traces into the provided store
func RunIngester(nsqLookupdHTTPAddrs []string, st Store) {

	cfg := nsq.NewConfig()
	cfg.UserAgent = fmt.Sprintf("phosphor go-nsq/%s", nsq.VERSION)
	cfg.MaxInFlight = maxInFlight

	consumer, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		log.Critical(err)
		os.Exit(1)
	}

	consumer.AddConcurrentHandlers(&IngestionHandler{
		store: st,
	}, 10)

	err = consumer.ConnectToNSQLookupds(nsqLookupdHTTPAddrs)
	if err != nil {
		log.Critical(err)
		os.Exit(1)
	}

	// Block until exit
	<-consumer.StopChan
}

// IngestionHandler exists to match the NSQ handler interface
type IngestionHandler struct {
	store Store
}

// HandleMessage delivered by NSQ
func (ih *IngestionHandler) HandleMessage(message *nsq.Message) error {

	p := &traceproto.Annotation{}
	err := proto.Unmarshal(message.Body, p)
	if err != nil {
		// returning an error to NSQ will requeue this
		// failure to unmarshal is permanent
		return nil
	}

	a := ProtoToAnnotation(p)
	log.Debugf("Received annotation: %+v", a)

	// Write to our store
	ih.store.StoreAnnotation(a)

	return nil
}
