package phosphor

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/mondough/phosphor/internal/util"
)

var HTTPPort = 7750

var nsqLookupdHTTPAddrs = util.StringArray{}

func init() {
	flag.Var(&nsqLookupdHTTPAddrs, "nsq-lookupd-http-address", "nsqlookupd HTTP address (may be given multiple times)")
}

type phosphor struct {
	Store Store
}

type PhosphorOptions struct {
	Store Store
}

func New(opts *PhosphorOptions) *phosphor {
	return &phosphor{
		Store: opts.Store,
	}
}

func (p *phosphor) Run() {

}

func main() {
	log.Infof("Phosphor starting up")
	defer log.Flush()

	flag.Parse()

	// Initialise a persistent store
	DefaultStore = NewMemoryStore()

	// Initialise trace ingestion
	go RunIngester(nsqLookupdHTTPAddrs, DefaultStore)

	// Set up API and serve requests
	http.HandleFunc("/", Index)
	http.HandleFunc("/trace", TraceLookup)
	http.ListenAndServe(fmt.Sprintf(":%v", HTTPPort), nil)
}
