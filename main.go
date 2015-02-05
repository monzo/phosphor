package main

import (
	"flag"

	log "github.com/cihub/seelog"

	"github.com/mattheath/phosphor/ingester"
	"github.com/mattheath/phosphor/memorystore"
	"github.com/mattheath/phosphor/util"
)

var APIPort = 7772

var nsqLookupdHTTPAddrs = util.StringArray{}

func init() {
	flag.Var(&nsqLookupdHTTPAddrs, "nsq-lookupd-http-address", "nsqlookupd HTTP address (may be given multiple times)")
}

func main() {
	defer log.Flush()

	flag.Parse()

	log.Infof("Phosphor starting up")

	store := memorystore.New()

	go ingester.Run(nsqLookupdHTTPAddrs, store)

}
