package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/users", h.List)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListAllUsers(r.Context())
	if err != nil {
		http.Error(w, "Erro ao buscar usuários do banco", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
