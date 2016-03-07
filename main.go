package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	log "github.com/cihub/seelog"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"goweb-scaffold/config"
	"goweb-scaffold/cron"
	"goweb-scaffold/rest"
)

var taskRunner cron.TaskRunner
var restHandler rest.RestHandler
var globalConfig config.GlobalConfig

func main() {
	// flag parsing
	var port string
	flag.StringVar(&port, "port", "8000", "port")
	flag.Parse()

	// config loading
	globalConfig = config.LoadGcloudConfig(
		config.LoadAsset("/config/config.json"))
	log.Debugf("config: %+v", globalConfig)

	buildDependencyGraph()

	// run cron job
	taskRunner.GlobalRun()
	defer cron.GlobalCron.Stop()

	// http server
	n := negroni.Classic()
	router := rest.BuildRouter(restHandler)
	n.Use(negroni.HandlerFunc(cors))
	n.UseHandler(router)
	n.Run(fmt.Sprintf(":%s", port))
}

// buildDependencyGraph builds dependency graph
func buildDependencyGraph() {
	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &globalConfig},
		&inject.Object{Value: &taskRunner},
		&inject.Object{Value: &restHandler},
	)
	if err != nil {
		os.Exit(1)
	}
	if err := g.Populate(); err != nil {
		os.Exit(1)
	}
	// :~)
}

// cors middleware (cross-origin resource sharing)
func cors(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-CSRF-Token, X-PINGOTHER")

	if req.Method == "OPTIONS" {
		method := req.Header.Get("Access-Control-Request-Method")

		if method == "" {
			http.Error(rw, "Bad Request", http.StatusBadRequest)
			return
		}

		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		return
	}

	next(rw, req)
}
