package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/sikemausa/micro-service-example/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email)
	if err != nil {
		fmt.Printf("Error creating user in db: %s", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User

	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return nil
}
