package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/trunov/erply-assignement-task/user-service/internal/repository/erply"
)

type UseCase interface {
	GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (erply.Customer, error)
}

type Handler struct {
	useCase    UseCase
	clientCode string
}

func New(useCase UseCase, clientCode string) *Handler {
	return &Handler{
		useCase:    useCase,
		clientCode: clientCode,
	}
}

func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	sessionKey := r.Context().Value("sessionKey").(string)

	id := chi.URLParam(r, "id")

	customer, err := h.useCase.GetCustomer(ctx, sessionKey, h.clientCode, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
