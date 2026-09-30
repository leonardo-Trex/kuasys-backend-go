package user

import (
	"context"

	"github.com/google/uuid"
	db "github.com/leonardo-Trex/kuasys-backend-go/internal/db/sqlc"
)

type Service struct {
	repo *db.Queries
}

func NewService(repo *db.Queries) *Service {
	return &Service{repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (db.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (s *Service) FindAll(ctx context.Context) ([]db.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// TODO entender isso
	if users == nil {
		return []db.User{}, nil
	}

	return users, nil
}

func (s *Service) DeleteByID(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
