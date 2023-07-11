package router

import (
	"github.com/go-chi/chi"
	"github.com/trunov/erply-assignement-task/user-service/internal/transport/handler"
)

func NewRouter(h *handler.Handler) chi.Router {
	r := chi.NewRouter()

	r.Get("/customer", h.GetCustomer)

	return r
}
