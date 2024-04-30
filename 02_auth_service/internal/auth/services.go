package auth

import (
	"context"
	"fmt"
	"log/slog"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error)
	GetByEmailAndPassword(ctx context.Context, loginUser LoginUserDTO) (*User, error)
}

type service struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	storage Storage
	logger  *slog.Logger
}

func NewService(repo *Repository, log *slog.Logger) *service {
	return &service{storage: repo, logger: log}
}

func (s *service) CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error) {
	s.logger.Debug("check password and repeat password")
	if dto.Password != dto.RepeatPassword {
		fmt.Println(err.Error())
	}

	user := NewUser(dto)

	s.logger.Debug("generate password hash")
	err = user.GeneratePasswordHash()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userUUID, err = s.storage.CreateUser(ctx, user)

	if err != nil {
		fmt.Println(err.Error())
	}

	return userUUID, nil
}

func (s *service) GetByEmailAndPassword(ctx context.Context, loginUser LoginUserDTO) (*User, error) {

	lgUser := LoginUser(loginUser)

	user, err := s.storage.GetByEmail(ctx, lgUser.Email)

	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lgUser.Password)); err != nil {
		return nil, err
	}

	return user, nil
}
