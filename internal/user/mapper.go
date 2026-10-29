package user

import (
	"uuid"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
)

// Os dados do usuário são divididos entre duas tabelas para permitir dois tipos de
// cadastro. Essa função pega o resultado da consulta com todos os dados e mapeia
// para o DTO.
func byIDtoResponse(e db.FindUserWithProfileByIDRow) (UserResponse, error) {
	response := UserResponse{
		uuid.UUID(e.ID),
		e.CreatedAt,
		e.UpdatedAt,
		e.KeycloakID,
		e.Email,
		e.Name,
		e.Cpf,
		e.Phone,
	}

	return response, nil
}

// O sqlc cria um tipo para as duas consultas, isso parece problemático.
// Esses dois tipos são na prática os mesmos
// esse erro tá meio perdido.
func findAlltoResponse(e db.FindAllRow) (UserResponse, error) {
	response := UserResponse{
		uuid.UUID(e.ID),
		e.CreatedAt,
		e.UpdatedAt,
		e.KeycloakID,
		e.Email,
		e.Name,
		e.Cpf,
		e.Phone,
	}

	return response, nil
}
