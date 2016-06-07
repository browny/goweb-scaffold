// Package rest represents the REST layer
package rest

import (
	"fmt"
	"net/http"

	"goweb-scaffold/config"
	"goweb-scaffold/logger"

	"github.com/gorilla/mux"
)

// Error is self defined error type
type Error struct {
	Error  error
	Detail string
	Code   int
}

// handlerWrapper manages all http error handling
type handlerWrapper func(http.ResponseWriter, *http.Request) *Error

func (fn handlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if restErr := fn(w, r); restErr != nil {
		if restErr.Detail != "" {
			logger.Errorf("error detail: %s", restErr.Detail)
		}
		http.Error(w, restErr.Error.Error(), restErr.Code)
	}
}

// Handler includes all http handler methods
type Handler struct {
	*config.AppContext `inject:""`
}

// HealthCheck is called by Google cloud to do health check
func (rest *Handler) HealthCheck(w http.ResponseWriter, req *http.Request) *Error {
	fmt.Fprintf(w, rest.AppContext.ProjectID+" is OK")
	return nil
}

// BuildRouter registers all routes
func BuildRouter(h Handler) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/healthcheck",
		handlerWrapper(h.HealthCheck)).Methods("GET")

	return router
}
