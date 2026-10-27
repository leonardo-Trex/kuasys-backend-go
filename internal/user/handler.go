package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	// r.Get("/users", h.FindAll)
	r.Get("/users/{id}", h.FindByID)
	// r.Post("/users", h.Create)
	// r.Put("/users", h.Update)
	// r.Delete("/users/{id}", h.DeleteByID)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user, err := h.service.FindByID(r.Context(), uuid.MustParse(id)) // TODO tratar o panic...
	if err != nil {
		http.Error(w, "Erro ao buscar no banco", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
// 	users, err := h.service.FindAll(r.Context())
// 	if err != nil {
// 		http.Error(w, "Erro ao buscar usuários do banco", http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(users)
// }

// func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")

// 	parsedID, err := uuid.Parse(id)
// 	if err != nil {
// 		http.Error(w, "ID inválido", http.StatusBadRequest)
// 		return
// 	}

// 	if err := h.service.DeleteByID(r.Context(), parsedID); err != nil {
// 		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }
