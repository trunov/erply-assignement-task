package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/trunov/erply-assignement-task/user-service/cmd/migrate"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/storage"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/handler"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/router"
	use_case "github.com/trunov/erply-assignement-task/user-service/internal/use-case"
)

type App struct {
	HttpServer      *http.Server
	ErplyHttpClient *http.Client
}

func New(cfg config.Config) (*App, error) {
	err := migrate.Migrate(cfg.DatabaseDSN, migrate.Migrations)
	if err != nil {
		return nil, err
	}

	repo, err := storage.New(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	erplyHttpClient := http.DefaultClient

	erplyClient := erply.New(erplyHttpClient, cfg.ClientCode)

	uc := use_case.New(repo, erplyClient)

	h := handler.New(uc, cfg.ClientCode)
	r := router.NewRouter(h, erplyClient, cfg)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	return &App{
		HttpServer:      s,
		ErplyHttpClient: erplyHttpClient,
	}, nil
}

func (a *App) Run() error {
	log.Printf("starting server")
	return a.HttpServer.ListenAndServe()
}
