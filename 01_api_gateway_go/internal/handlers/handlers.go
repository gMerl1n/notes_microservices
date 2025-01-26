package handlers

import (
	"github.com/gMerl1n/notes_microservices/internal/clients/auth_client"
	"github.com/gMerl1n/notes_microservices/internal/handlers/auth_handlers"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator"
)

type Handlers struct {
	HandlersUser *auth_handlers.HandlerUser
}

func NewHandlers(
	clientUser auth_client.IClientUser,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	logger *logging.Logger) *Handlers {

	return &Handlers{
		HandlersUser: auth_handlers.NewHandlerUser(clientUser, jwtParser, validator, logger),
	}
}
