package address

import (
	"uuid"

	db "github.com/leonardo-Trex/kuasys-backend-go/internal/database/sqlc"
	"github.com/leonardo-Trex/kuasys-backend-go/internal/user"
)

func AddressByIDToResponse(addr db.FindAddressWithUserDetailsByIDRow) AddressResponse {
	return AddressResponse{
		uuid.UUID(addr.ID),
		addr.CreatedAt,
		addr.UpdatedAt,
		user.UserResponseSimple{
			ID:    uuid.UUID(addr.UserID),
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
