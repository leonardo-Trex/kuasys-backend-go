package address

import (
	"context"

	"github.com/google/uuid"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
)

type Service struct {
	repo *db.Queries
}

func NewService(repo *db.Queries) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (AddressResponse, error) {
	addr, err := s.repo.FindAddressWithUserDetailsByID(ctx, id)
	if err != nil {
		return AddressResponse{}, err
	}

	//TODO tratar o erro, mas não há valores null
	response := AddressByIDToResponse(addr)
	return response, nil
}
