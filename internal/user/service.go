package user

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

// Retorna tudo sobre o usuário
func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (UserResponse, error) {
	user, err := s.repo.FindUserWithProfileByID(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}

	//TODO tratar o erro, mas não há valores null
	response, _ := toResponse(user)
	return response, nil
}

// func (s *Service) FindAll(ctx context.Context) ([]db.User, error) {
// 	users, err := s.repo.FindAll(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO entender isso
// 	if users == nil {
// 		return []db.User{}, nil
// 	}

// 	return users, nil
// }

// func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
// 	err := s.repo.DeleteUser(ctx, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
