package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sikemausa/micro-service-example/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return status.Error(codes.Internal, "not found")
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
			return domain.User{}, status.Error(codes.NotFound, "not found")
		}
		return domain.User{}, status.Error(codes.Internal, "internal server error")
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	if err != nil {
		fmt.Printf("Error updating user in db: %s", err)
		return domain.User{}, status.Error(codes.Internal, "internal server error")
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Printf("Error deleting user in db: %s", err)
		return status.Error(codes.Internal, "internal server error")
	}

	return nil
}

func (r *UserRepository) List(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	query := `SELECT id, name, email FROM users`
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Printf("Error listing users from db: %s", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			fmt.Printf("Error scanning user row: %s", err)
			return nil, status.Error(codes.Internal, "internal server error")
		}
		users = append(users, user)
	}

	return users, nil
}
