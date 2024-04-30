package auth

import (
	"context"
	"fmt"
	"log/slog"
)

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) (userUUID string, err error)
	GetUserByUUID(ctx context.Context, uuid string) (*User, error)
}

type service struct {
	// создаем структуру, которая принимает репозиторий для работы с БД
	storage Storage
	logger  slog.Logger
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

func (s *service) GetUserByUUID(ctx context.Context, uuid string) (*User, error) {
	return nil, nil
}
