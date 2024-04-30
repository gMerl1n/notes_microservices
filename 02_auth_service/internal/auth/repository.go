package auth

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type AuthRepository struct {
// }

const (
	usersTable = "users"
)

type Storage interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type Repository struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewRepository(db *pgxpool.Pool, logger *slog.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (s *Repository) CreateUser(ctx context.Context, user User) (string, error) {

	var newUserUUID string

	query := fmt.Sprintf("INSERT INTO %s (UUID, email, password_hash) VALUES ($1, $2, $3) RETURNING UUID", usersTable)

	if err := s.db.QueryRow(ctx, query,
		user.UUID,
		user.Email,
		user.Password,
	).Scan(&newUserUUID); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return newUserUUID, nil
}

func (s *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {

	var user User

	query := fmt.Sprintf("SELECT UUID, email, password_hash FROM %s WHERE email=$1", usersTable)

	if err := s.db.QueryRow(ctx, query,
		email,
	).Scan(&user.UUID, &user.Email, &user.Password); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil

}
