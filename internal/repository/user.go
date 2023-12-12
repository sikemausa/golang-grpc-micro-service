package repository

import (
	"context"

	"github.com/sikemausa/micro-service-example/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user domain.User) error
	GetByID(ctx context.Context, id string) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]domain.User, error)
}
