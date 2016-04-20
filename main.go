package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"goweb-scaffold/config"
	"goweb-scaffold/cron"
	"goweb-scaffold/rest"

	log "github.com/cihub/seelog"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
)

var taskRunner cron.TaskRunner
var restHandler rest.RestHandler
var globalConfig config.GlobalConfig

func main() {
	// flag parsing
	var env string
	var port string
	flag.StringVar(&env, "env", "Alpha", "environment")
	flag.StringVar(&port, "port", "8000", "port")
	flag.Parse()

	// set logger
	seelogConf, _ := ioutil.ReadAll(config.LoadAsset("/config/seelog.xml"))
	logger, _ := log.LoggerFromConfigAsBytes(seelogConf)
	log.ReplaceLogger(logger)

	// config loading
	globalConfig = config.LoadGcloudConfig(
		config.LoadAsset("/config/config.json"))

	buildDependencyGraph()

	log.Debugf("App starts: env[%s], projectID[%s]", env, globalConfig.ProjectId)

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
