// Package rest represents the REST layer
package rest

import (
	"fmt"
	"net/http"

	"goweb-scaffold/config"
	"goweb-scaffold/logger"

	"github.com/gorilla/mux"
)

type restError struct {
	Error  error
	Detail string
	Code   int
}

// handlerWrapper manages all http error handling
type handlerWrapper func(http.ResponseWriter, *http.Request) *restError

func (fn handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if restErr := fn(w, r); restErr != nil {
		if restErr.Detail != "" {
			logger.Errorf("error detail: %s", restErr.Detail)
		}
		http.Error(w, restErr.Error.Error(), restErr.Code)
	}
}

// RestHandler includes all http handler methods
type RestHandler struct {
	*config.AppContext `inject:""`
}

// HealthCheck is called by Google cloud to do health check
func (rest *RestHandler) HealthCheck(w http.ResponseWriter, req *http.Request) *restError {
	fmt.Fprintf(w, rest.AppContext.ProjectID+" is OK")
	return nil
}

func BuildRouter(restHandler RestHandler) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/healthcheck",
		handlerWrapper(restHandler.HealthCheck)).Methods("GET")

	return router
}
