package handlers

import (
	"github.com/gMerl1n/notes_microservices/internal/clients"
	"github.com/gMerl1n/notes_microservices/internal/clients/notes_server_clients"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	clientUser       clients.IClientUser
	clientNotes      notes_server_clients.IClientNotes
	clientCategories notes_server_clients.IClientCategories
	jwtParser        jwt.ITokenParser
	validator        *validator.Validate
	logger           *logging.Logger
}

func NewHandler(
	clientUser clients.IClientUser,
	clientNotes notes_server_clients.IClientNotes,
	clientCategories notes_server_clients.IClientCategories,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	logger *logging.Logger) *Handler {

	return &Handler{
		clientUser:       clientUser,
		clientNotes:      clientNotes,
		clientCategories: clientCategories,
		jwtParser:        jwtParser,
		validator:        validator,
		logger:           logger,
	}
}

// type Handlers struct {
// 	HandlersUser      *HandlerUser
// 	HandlersNotes     *HandlerNotes
// 	HandlerCategories *HandlerCategories
// }

// func NewHandlers(
// 	clientUser clients.IClientUser,
// 	clientNotes notes_server_clients.IClientNotes,
// 	clientCategories notes_server_clients.IClientCategories,
// 	jwtParser jwt.ITokenParser,
// 	validator *validator.Validate,
// 	logger *logging.Logger) *Handlers {

// 	return &Handlers{
// 		HandlersUser:      NewHandlerUser(clientUser, jwtParser, validator, logger),
// 		HandlersNotes:     NewHandlerNotes(clientNotes, jwtParser, validator, logger),
// 		HandlerCategories: NewHandlerCategories(clientCategories, jwtParser, validator, logger),
// 	}
// }
