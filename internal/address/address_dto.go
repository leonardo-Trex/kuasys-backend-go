package address

import (
	"github.com/google/uuid"

	"github.com/leonardo-Trex/kuasys-backend-go/internal/user"
)

type AddressResponse struct {
	ID           uuid.UUID               `json:"id"`
	User         user.UserResponseSimple `json:"user"`
	Cep          string                  `json:"cep"`
	Number       string                  `json:"number"`
	Street       string                  `json:"street"`
	Complement   string                  `json:"complement"`
	Neighborhood string                  `json:"neighborhood"`
	City         string                  `json:"city"`
	State        string                  `json:"state"`
	IsPrimary    bool                    `json:"is_primary"`
}

type AddressCreateRequest struct {
	UserID       uuid.UUID `json:"user_id"`
	Cep          string    `json:"cep"`
	Number       string    `json:"number"`
	Street       string    `json:"street"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	IsPrimary    bool      `json:"is_primary"`
}

type AddressCreateResponse struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Cep          string    `json:"cep"`
	Number       string    `json:"number"`
	Street       string    `json:"street"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	IsPrimary    bool      `json:"is_primary"`
}

type AddressUpdateRequest struct {
	Cep          string `json:"cep"`
	Number       string `json:"number"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	IsPrimary    bool   `json:"is_primary"`
}

type AddressUpdateResponse struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	Cep          string    `json:"cep"`
	Number       string    `json:"number"`
	Street       string    `json:"street"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	IsPrimary    bool      `json:"is_primary"`
}
