// Package rest represents the REST layer
package rest

import (
	"fmt"
	"net/http"

	log "github.com/cihub/seelog"
	"github.com/gorilla/mux"
)

type restError struct {
	Error  error
	Detail string
	Code   int
}

// RestHandlerWrapper manages all http error handling
type RestHandlerWrapper func(http.ResponseWriter, *http.Request) *restError

func (fn RestHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if restErr := fn(w, r); restErr != nil {
		if restErr.Detail != "" {
			log.Errorf("error detail: %s", restErr.Detail)
		}
		http.Error(w, restErr.Error.Error(), restErr.Code)
	}
}

// RestHandler includes all http handler methods
type RestHandler struct {
}

// HealthCheck is called by Google cloud to do health check
func (rest *RestHandler) HealthCheck(w http.ResponseWriter, req *http.Request) *restError {
	fmt.Fprintf(w, "OK World")
	return nil
}

func BuildRouter(restHandler RestHandler) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/healthcheck",
		RestHandlerWrapper(restHandler.HealthCheck)).Methods("GET")

	return router
}
