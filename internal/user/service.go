package user

import (
	"context"

	"github.com/leonardo-Trex/kuasys-backend-go/internal/db"
)

type Service struct {
	repo *db.Queries
}

func NewService(repo *db.Queries) *Service {
	return &Service{repo: repo}
}
func (s *Service) ListAllUsers(ctx context.Context) ([]db.User, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return []db.User{}, nil
	}

	return users, nil
}
