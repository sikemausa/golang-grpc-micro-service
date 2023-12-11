package service

import (
	"context"
	"github.com/sikemausa/micro-service-example/internal/repository"

	"github.com/sikemausa/micro-service-example/internal/domain"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, user domain.User) (domain.User, error) {
    err := s.repo.Create(ctx, user)
    if err != nil {
        return domain.User{}, err
    }
    return user, nil
}

func (s *UserService) Get(ctx context.Context, id string) (domain.User, error) {
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return domain.User{}, err
    }
    return user, nil
}

func (s *UserService) Update(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	return nil
}
