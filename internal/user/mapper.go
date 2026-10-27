package user

import (
	"uuid"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
)

func toResponse(e db.FindUserWithProfileByIDRow) (UserResponse, error) {
	response := UserResponse{
		uuid.UUID(e.ID),
		e.CreatedAt,
		e.UpdatedAt,
		e.KeycloakID,
		e.Email,
		e.IsActive,
		e.Name,
		e.Cpf,
		e.Phone,
	}

	return response, nil
}
