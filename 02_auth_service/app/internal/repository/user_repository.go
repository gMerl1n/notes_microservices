package repository

import (
	"context"
	"fmt"

	"github.com/gMerl1n/notes_microservices/app/internal/config"
	"github.com/gMerl1n/notes_microservices/app/internal/domain"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	usersTable = "users"
)

type IRepositoryUser interface {
	CreateUser(ctx context.Context, name, surname, email, hashedPassword string, age int) (int, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}

type RepositoryUser struct {
	db         *pgxpool.Pool
	logger     *logging.Logger
	configUser *config.ConfigUser
}

func NewRepositoryUser(db *pgxpool.Pool, logger *logging.Logger, configUser *config.ConfigUser) *RepositoryUser {
	return &RepositoryUser{db: db, logger: logger, configUser: configUser}
}

func (s *RepositoryUser) CreateUser(ctx context.Context, name, surname, email, hashedPassword string, age int) (int, error) {

	var userID int

	query := fmt.Sprintf(
		`INSERT INTO %s (name, surname, age, email, password_hash, role_id) 
		 VALUES ($1, $2, $3, $4, $5, $6) 
		 RETURNING id`,
		usersTable,
	)

	if err := s.db.QueryRow(ctx, query,
		name,
		surname,
		age,
		email,
		hashedPassword,
		s.configUser.UserRoleID,
	).Scan(&userID); err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return userID, nil
}

func (s *RepositoryUser) GetByEmail(ctx context.Context, email string) (*domain.User, error) {

	var user domain.User

	query := fmt.Sprintf(
		`SELECT id, email, password_hash 
		 FROM %s 
		 WHERE email=$1`,
		usersTable)

	if err := s.db.QueryRow(ctx, query,
		email,
	).Scan(&user.ID, &user.Email, &user.Password); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil

}
