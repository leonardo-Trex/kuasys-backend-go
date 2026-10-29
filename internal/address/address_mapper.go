package address

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
	"github.com/leonardo-Trex/kuasys-backend-go/internal/user"
)

func AddressByIDToResponse(addr db.FindAddressWithUserDetailsByIDRow) AddressResponse {
	return AddressResponse{
		uuid.UUID(addr.ID),
		user.UserResponseSimple{
			ID:    addr.UserID,
			Name:  addr.UserName,
			Email: addr.UserEmail,
			Cpf:   addr.UserCpf,
			Phone: addr.UserPhone,
		},
		addr.Cep,
		addr.Number,
		addr.Street,
		addr.Complement.String,
		addr.Neighborhood,
		addr.City,
		addr.State,
		addr.IsPrimary,
	}
}

func AddressFindAllToResponse(addr db.FindAllAddressWithUserDetailsRow) AddressResponse {
	return AddressResponse{
		ID: addr.ID,
		User: user.UserResponseSimple{
			ID:    uuid.UUID(addr.UserID),
			Name:  addr.UserName,
			Email: addr.UserEmail,
			Cpf:   addr.UserCpf,
			Phone: addr.UserPhone,
		},
		Cep:          addr.Cep,
		Number:       addr.Number,
		Street:       addr.Street,
		Complement:   addr.Complement.String,
		Neighborhood: addr.Neighborhood,
		City:         addr.City,
		State:        addr.State,
		IsPrimary:    addr.IsPrimary,
	}
}

// O id de usuario deve ser controlado futuramente pelo JWT
// Um mapper deve ser estritamente deterministico
func CreateRequestToEntity(addr AddressCreateRequest, id uuid.UUID) db.CreateAddressParams {
	return db.CreateAddressParams{
		ID:     id,
		UserID: addr.UserID,
		Cep:    addr.Cep,
		Number: addr.Number,
		Street: addr.Street,
		// TODO entender strings, text e campos null e as consequencias de usar *string
		Complement: pgtype.Text{
			String: addr.Complement,
			Valid:  true,
		},
		Neighborhood: addr.Neighborhood,
		City:         addr.City,
		State:        addr.State,
		IsPrimary:    addr.IsPrimary,
	}
}

func CreateAddressToResponse(e db.Address) AddressCreateResponse {
	return AddressCreateResponse{
		ID:           e.ID,
		UserID:       e.UserID,
		Cep:          e.Cep,
		Number:       e.Number,
		Street:       e.Street,
		Complement:   e.Complement.String,
		Neighborhood: e.Neighborhood,
		City:         e.City,
		State:        e.State,
		IsPrimary:    e.IsPrimary,
	}
}

func UpdateAddressRequestToUpdateAddressParams(dto AddressUpdateRequest, uuid uuid.UUID) db.UpdateAddressParams {
	return db.UpdateAddressParams{
		ID:     uuid,
		Cep:    dto.Cep,
		Number: dto.Number,
		Street: dto.Street,
		//TODO consertar essa bagunça
		Complement: pgtype.Text{
			String: dto.Cep,
			Valid:  true,
		},
		Neighborhood: dto.Neighborhood,
		City:         dto.City,
		State:        dto.State,
		IsPrimary:    dto.IsPrimary,
	}
}

func AddressToUpdateAddressResponse(a db.Address) AddressUpdateResponse {
	return AddressUpdateResponse{
		ID:           a.ID,
		UserID:       a.UserID,
		Cep:          a.Cep,
		Number:       a.Number,
		Street:       a.Street,
		Complement:   a.Complement.String,
		Neighborhood: a.Neighborhood,
		City:         a.City,
		State:        a.State,
		IsPrimary:    a.IsPrimary,
	}
}
