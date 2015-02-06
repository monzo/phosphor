package handler

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm Phosphor")
}

func TraceLookup(w http.ResponseWriter, r *http.Request) {
	log.Infof("Trace Lookup - TraceId: %s", r.URL.Query().Get("traceId"))
}
