package user

import (
	"uuid"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
)

// Os dados do usuário são divididos entre duas tabelas para permitir dois tipos de
// cadastro. Essa função pega o resultado da consulta com todos os dados e mapeia
// para o DTO.
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
