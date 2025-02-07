package handlers

import (
	"github.com/gMerl1n/notes_microservices/internal/clients"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	HandlersUser *HandlerUser
}

func NewHandlers(
	clientUser clients.IClientUser,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	logger *logging.Logger) *Handlers {

	return &Handlers{
		HandlersUser: NewHandlerUser(clientUser, jwtParser, validator, logger),
	}
}
