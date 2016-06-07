package rest_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"goweb-scaffold/config"
	"goweb-scaffold/logger"

	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/stretchr/testify/suite"
	"goweb-scaffold/rest"
)

var n *negroni.Negroni
var appContext config.AppContext
var testedRestHandler rest.RestHandler

func TestRestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RestHandlerTestSuite))
}

type RestHandlerTestSuite struct {
	suite.Suite
}

func (suite *RestHandlerTestSuite) SetupSuite() {
	// flag parsing
	var env string
	flag.StringVar(&env, "env", "Alpha", "environment")
	flag.Parse()

	buildDependencyGraph(env)
	logger.SetupLogger()

	// run test http server
	n = negroni.New()
	router := rest.BuildRouter(testedRestHandler)
	n.UseHandler(router)

	logger.Debug("======== RestHandler Test Begin ========")
}

func (suite *RestHandlerTestSuite) TestHealthCheck() {
	t := suite.T()

	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	req, err := http.NewRequest("GET", "http://localhost/healthcheck", nil)
	if err != nil {
		t.Error(err)
	}
	n.ServeHTTP(response, req)

	if body, err := ioutil.ReadAll(response.Body); err != nil {
		t.Error(err)
	} else if string(body) != "ball is OK" {
		t.Error("expected", "ball is OK", "got", string(body))
	}
}

func buildDependencyGraph(env string) {
	config.Viper()
	appContext.Load(env)

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &appContext},
		&inject.Object{Value: &testedRestHandler},
	)
	if err != nil {
		os.Exit(1)
	}
	if err := g.Populate(); err != nil {
		os.Exit(1)
	}
	// :~)
}

func (suite *RestHandlerTestSuite) TearDownSuite() {
	logger.Debugf("======== RestHandler Test End ========")
}
