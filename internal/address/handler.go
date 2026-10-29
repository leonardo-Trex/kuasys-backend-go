package address

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	// r.Get("/users", h.FindAll)
	r.Get("/addresses/{id}", h.FindByID)
	// r.Post("/users", h.Create)
	// r.Put("/users", h.Update)
	// r.Delete("/users/{id}", h.DeleteByID)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	// Se o banco não tem registro user talvez possa ser nil? ou o erro tráz essa informação?
	user, err := h.service.FindByID(r.Context(), uuid.MustParse(id)) // TODO tratar o panic...
	//TODO isso tá estranho
	if err != nil {
		// TODO retorna o json do erro RFC 7807
		http.Error(w, "Erro ao buscar no banco", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
