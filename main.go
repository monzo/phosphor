package main

import (
	"flag"
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/mattheath/phosphor/handler"
	"github.com/mattheath/phosphor/ingester"
	"github.com/mattheath/phosphor/store"
	"github.com/mattheath/phosphor/util"
)

var HTTPPort = 7750

var nsqLookupdHTTPAddrs = util.StringArray{}

func init() {
	flag.Var(&nsqLookupdHTTPAddrs, "nsq-lookupd-http-address", "nsqlookupd HTTP address (may be given multiple times)")
}

func main() {
	log.Infof("Phosphor starting up")
	defer log.Flush()

	flag.Parse()

	// Initialise a persistent store
	store := store.NewMemoryStore()

	// Initialise trace ingestion
	go ingester.Run(nsqLookupdHTTPAddrs, store)

	// Set up API and serve requests
	handler.DefaultStore = store
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/trace", handler.TraceLookup)
	http.ListenAndServe(fmt.Sprintf(":%v", HTTPPort), nil)
}
