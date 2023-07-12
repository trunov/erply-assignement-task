package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-resty/resty/v2"
)

type UseCase interface {
	GetCustomer(ctx context.Context, sessionKey, clientCode, customerID string) (*resty.Response, error)
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
	sessionKey := "xxx"

	id := chi.URLParam(r, "id")

	resp, err := h.useCase.GetCustomer(ctx, sessionKey, h.clientCode, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(resp)
}
