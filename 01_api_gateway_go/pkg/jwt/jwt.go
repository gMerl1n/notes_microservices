package jwt

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// TokenManager provides logic for JWT & Refresh tokens generation and parsing.
type ITokenParser interface {
	Parse(accessToken string) (string, error)
}

type TokenParser struct {
	signingKey string
}

func NewManager(signingKey string) (*TokenParser, error) {
	if signingKey == "" {
		return nil, errors.New("empty signing key")
	}

	return &TokenParser{
		signingKey: signingKey}, nil
}

func (m *TokenParser) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}
