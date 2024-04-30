package auth

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type AuthRepository struct {
// }

type Storage interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUserByUUID(ctx context.Context, uuid string) (*User, error)
}

type Repository struct {
	db     *pgxpool.Pool
	logger slog.Logger
}

func NewRepository(db *pgxpool.Pool, logger slog.Logger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (s *Repository) CreateUser(ctx context.Context, user User) (string, error) {
	return "", nil
}

func (s *Repository) GetUserByUUID(ctx context.Context, uuid string) (*User, error) {
	return nil, nil
}
