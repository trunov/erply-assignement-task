package router

import (
	"github.com/go-chi/chi"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
	"github.com/trunov/erply-assignement-task/user-service/internal/middleware"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/handler"
	use_case "github.com/trunov/erply-assignement-task/user-service/internal/use-case"
)

func NewRouter(h *handler.Handler, erply use_case.Erply, cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.VerifyErplyUser(erply, cfg.ClientCode, cfg.Username, cfg.Password))
	r.Use(middleware.TokenAuthorization(cfg.Auth))

	r.Get("/customer/{id}", h.GetCustomer)

	return r
}
