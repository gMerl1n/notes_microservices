package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(uuid string, email string, duration time.Duration) (string, error) {

	Secret := "123"

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uuid
	claims["email"] = email
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
