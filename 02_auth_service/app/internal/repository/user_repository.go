package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1n/notes_microservices/app/internal/domain"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	usersTable = "users"
)

type IRepositoryUser interface {
	CreateUser(ctx context.Context, name, surname, email, hashedPassword string, age int) (string, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type RepositoryUser struct {
	db     *pgxpool.Pool
	logger *logging.Logger
}

func NewRepositoryUser(db *pgxpool.Pool, logger *logging.Logger) *RepositoryUser {
	return &RepositoryUser{db: db, logger: logger}
}

func (s *RepositoryUser) CreateUser(ctx context.Context, name, surname, email, hashedPassword string, age int) (string, error) {

	var newUserUUID string

	query := fmt.Sprintf(
		`INSERT INTO %s (name, surname, age, email, password_hash) 
		 VALUES ($1, $2, $3, $4, $5) 
		 RETURNING UUID`,
		usersTable,
	)

	if err := s.db.QueryRow(ctx, query,
		name,
		surname,
		age,
		email,
		hashedPassword,
	).Scan(&newUserUUID); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return newUserUUID, nil
}

func (s *RepositoryUser) GetByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	query := fmt.Sprintf(
		`SELECT UUID, email, password_hash 
		 FROM %s 
		 WHERE email=$1`,
		usersTable)

	if err := s.db.QueryRow(ctx, query,
		email,
	).Scan(&user.UUID, &user.Email, &user.Password); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil

}
