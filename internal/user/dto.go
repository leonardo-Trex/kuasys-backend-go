package user

import (
	"time"
	"uuid"
)

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	KeycloakId string    `json:"keycloak_id"`
	Email      string    `json:"email"`
	IsActive   bool      `json:"is_active"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	Phone      string    `json:"phone"`
}
