package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type UseCase interface {
	GetCustomer(ctx context.Context, id int)
}

type Handler struct {
	useCase UseCase
}

func New(useCase UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	h.useCase.GetCustomer(ctx, id)
}
