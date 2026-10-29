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
	r.Get("/addresses", h.FindAll)
	r.Get("/addresses/{id}", h.FindByID)
	r.Post("/addresses", h.Create)
	r.Put("/addresses/{id}", h.Update)
	r.Delete("/addresses/{id}", h.DeleteByID)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, _ := uuid.Parse(id) // TODO tratar o erro...

	addr, err := h.service.FindByID(r.Context(), parsedID)
	if err != nil {
		// TODO retorna o json do erro RFC 7807
		http.Error(w, "Erro ao buscar no banco", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addr)
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	// Se o banco não tem registro user talvez possa ser nil? ou o erro tráz essa informação?
	addrs, err := h.service.FindAll(r.Context()) // TODO tratar o panic...
	//TODO isso tá estranho
	if err != nil {
		// TODO retorna o json do erro RFC 7807
		http.Error(w, "Erro ao buscar no banco", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(addrs)
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteByID(r.Context(), parsedID); err != nil {
		// TODO retorna o json do erro RFC 7807
		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// TODO usar ponteiros tras o valor nil que pode ser util.
// Não estrou tratando erro nenhum...
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddressCreateRequest

	// TODO criar meios de validação
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	//TODO tratar o erro
	resp, _ := h.service.CreateAddress(r.Context(), req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	parsedID, _ := uuid.Parse(id)

	var req AddressUpdateRequest

	// TODO criar meios de validação
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// TODO retorna o json do erro RFC 7807
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}

	//TODO tratar o erro
	resp, _ := h.service.UpdateAddress(r.Context(), req, parsedID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) // TODO Preciso tratar o erro aqui?
}
