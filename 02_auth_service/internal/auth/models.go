package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID     string
	Email    string
	Password string
}

type CreateUserDTO struct {
	Email          string
	Password       string
	RepeatPassword string
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return fmt.Errorf("password does not match")
	}
	return nil
}

func (u *User) GeneratePasswordHash() error {
	pwd, err := generatePasswordHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = pwd
	return nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password due to error %w", err)
	}
	return string(hash), nil
}
