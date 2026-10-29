package user

import (
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	KeycloakId string    `json:"keycloak_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	Phone      string    `json:"phone"`
}

type UserResponseSimple struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Cpf   string    `json:"cpf"`
	Phone string    `json:"phone"`
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Phone string `json:"phone"`
}
