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

func (s *Service) FindAll(ctx context.Context) ([]AddressResponse, error) {
	addrs, err := s.repo.FindAllAddressWithUserDetails(ctx)
	if err != nil {
		return []AddressResponse{}, err
	}

	var response []AddressResponse
	for _, addr := range addrs {
		response = append(response, AddressFindAllToResponse(addr))
	}

	return response, nil
}

// Eu ainda não entendi onde entra as transações
func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteAddress(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Eu ainda não entendi onde entra as transações
func (s *Service) CreateAddress(ctx context.Context, dto AddressCreateRequest) (AddressCreateResponse, error) {
	// TODO Tratar o erro
	uuid, _ := uuid.NewV7()

	addr, err := s.repo.CreateAddress(ctx, CreateRequestToEntity(dto, uuid))

	// Eu não devo retornar o zero value da struct  aqui.
	if err != nil {
		return AddressCreateResponse{}, err
	}

	return CreateAddressToResponse(addr), nil
}

// Eu ainda não entendi onde entra as transações
func (s *Service) UpdateAddress(ctx context.Context, dto AddressUpdateRequest, id uuid.UUID) (AddressUpdateResponse, error) {

	addr, err := s.repo.UpdateAddress(ctx, UpdateAddressRequestToUpdateAddressParams(dto, id))

	// Eu não devo retornar o zero value da struct  aqui.
	if err != nil {
		return AddressUpdateResponse{}, err
	}

	resp := AddressToUpdateAddressResponse(addr)

	return resp, nil
}
