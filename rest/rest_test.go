package rest_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	log "github.com/cihub/seelog"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/suite"
	"goweb-scaffold/rest"
)

var n *negroni.Negroni
var testedRestHandler rest.RestHandler

func TestRestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RestHandlerTestSuite))
}

type RestHandlerTestSuite struct {
	suite.Suite
}

func (suite *RestHandlerTestSuite) SetupSuite() {
	n = negroni.New()
	router := rest.BuildRouter(testedRestHandler)
	n.UseHandler(router)

	log.Debug("======== RestHandler Test Begin ========")
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
	} else if string(body) != "OK" {
		t.Error("expected", "OK", "got", body)
	}
}

func (suite *RestHandlerTestSuite) TearDownSuite() {
	log.Debugf("======== RestHandler Test End ========")
}
