package address

import (
	"time"
	"uuid"

	"github.com/leonardo-Trex/kuasys-backend-go/internal/user"
)

type AddressResponse struct {
	ID           uuid.UUID               `json:"id"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
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
