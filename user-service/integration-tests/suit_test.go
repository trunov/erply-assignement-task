package integrationtests

import (
	"context"
	"database/sql"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"github.com/trunov/erply-assignement-task/testutils"
	"github.com/trunov/erply-assignement-task/testutils/testcontainer"
	"github.com/trunov/erply-assignement-task/user-service/internal/app"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
)

var (
	ClientCode = "532389"
	Auth       = "admin"
	Username   = "john_doe"
	Password   = "s3cr3t"
)

type TestSuite struct {
	suite.Suite
	app               *app.App
	server            *httptest.Server
	postgresContainer *testcontainer.PostgresContainer
}

func (s *TestSuite) SetupSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer ctxCancel()

	var err error

	s.postgresContainer, err = testcontainer.NewPostgresContainer(ctx)
	s.Require().NoError(err)

	s.app, err = app.New(config.Config{
		DatabaseDSN: s.postgresContainer.GetDSN(),
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

func (s *TestSuite) SetupTest() {
	db, err := sql.Open("postgres", s.postgresContainer.GetDSN())
	s.Require().NoError(err)

	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.FS(testutils.Fixtures),
		testfixtures.Directory("fixtures/storage"),
	)

	s.Require().NoError(err)
	s.Require().NoError(fixtures.Load())
}

func (s *TestSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	httpmock.DeactivateAndReset()

	err := s.postgresContainer.Terminate(ctx)
	s.Require().NoError(err)
}

func TestSuite_Run(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
