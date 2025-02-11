package handlers

import (
	"github.com/gMerl1n/notes_microservices/internal/clients/notes_server_clients"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
)

type HandlerCategories struct {
	clientCategories notes_server_clients.IClientCategories
	jwtParser        jwt.ITokenParser
	validator        *validator.Validate
	logger           *logging.Logger
}

func NewHandlerCategories(
	clientCategories notes_server_clients.IClientCategories,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	log *logging.Logger,

) *HandlerCategories {
	return &HandlerCategories{
		clientCategories: clientCategories,
		jwtParser:        jwtParser,
		validator:        validator,
		logger:           log,
	}
}
