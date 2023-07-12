package router

import (
	"github.com/go-chi/chi"
	"github.com/trunov/erply-assignement-task/user-service/internal/config"
	"github.com/trunov/erply-assignement-task/user-service/internal/middleware"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/handler"
)

func NewRouter(h *handler.Handler, cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.TokenAuthorization(cfg.Auth))

	r.Get("/customer/{id}", h.GetCustomer)

	return r
}
