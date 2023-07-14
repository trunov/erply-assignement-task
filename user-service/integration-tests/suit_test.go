package integrationtests

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"github.com/trunov/erply-assignement-task/user-service/internal/app"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
)

const DatabaseDSN = "postgres://....."

var (
	ClientCode = "532389"
	Auth       = "admin"
	Username   = "john_doe"
	Password   = "s3cr3t"
)

type TestSuite struct {
	suite.Suite
	app    *app.App
	server *httptest.Server
}

func (s *TestSuite) SetupSuite() {
	_, ctxCancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer ctxCancel()

	var err error
	s.app, err = app.New(config.Config{
		DatabaseDSN: DatabaseDSN,
		Port:        8080,
		ClientCode:  ClientCode,
		Username:    Username,
		Password:    Password,
		Auth:        Auth,
	})
	s.Require().NoError(err)

	s.server = httptest.NewServer(s.app.HttpServer.Handler)

	httpmock.ActivateNonDefault(s.app.ErplyHttpClient)
}

func (s *TestSuite) TearDownSuite() {
	_, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	httpmock.DeactivateAndReset()
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
