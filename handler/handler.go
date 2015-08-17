package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"

	"github.com/mondough/phosphor/store"
	"github.com/mondough/phosphor/util"
)

// DefaultStore is a reference to our persistence layer which we can query
var DefaultStore store.Store

// Index
// @todo return version information etc
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf(`{
		"name": "phosphor",
		"version": "%s"
	}`, util.VERSION))
}

// TraceLookup retrieves a trace from the persistence layer
func TraceLookup(w http.ResponseWriter, r *http.Request) {
	traceId := r.URL.Query().Get("traceId")
	if traceId == "" {
		errorResponse(r, w, http.StatusBadRequest, errors.New("traceId param not provided"))
		return
	}

	log.Debugf("Trace lookup - TraceId: %s", traceId)
	t, err := DefaultStore.ReadTrace(traceId)
	if err != nil {
		log.Errorf("Trace lookup failed: %s", err)
		errorResponse(r, w, http.StatusInternalServerError, fmt.Errorf("could not load trace: %s", err))
		return
	}

	// If we don't find the trace return 404
	if t == nil {
		log.Debugf("Trace not found: %s", traceId)
		errorResponse(r, w, http.StatusNotFound, errors.New("traceId not found"))
		return
	}

	// Return trace
	response(
		r,
		w,
		map[string]interface{}{
			"trace": prettyFormatTrace(t),
		},
	)
}

// response sends the response back to the client, marshaling to JSON
func response(r *http.Request, w http.ResponseWriter, resp interface{}) {
	writeResponse(r, w, http.StatusOK, resp)
}

// errorResponse marshals an error to JSON and returns this to the client
func errorResponse(r *http.Request, w http.ResponseWriter, code int, err error) {
	resp := map[string]interface{}{
		"error": err.Error(),
	}

	writeResponse(r, w, code, resp)
}

// response marshals a response to json and returns to the client
func writeResponse(r *http.Request, w http.ResponseWriter, code int, resp interface{}) {

	// Deal with CORS
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, POST, PUT")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// Allow any headers
		if wantedHeaders := r.Header.Get("Access-Control-Request-Headers"); wantedHeaders != "" {
			w.Header().Set("Access-Control-Allow-Headers", wantedHeaders)
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error":"failed to marshal json"}`)
		return
	}

	w.WriteHeader(code)
	fmt.Fprintln(w, string(b))
}
