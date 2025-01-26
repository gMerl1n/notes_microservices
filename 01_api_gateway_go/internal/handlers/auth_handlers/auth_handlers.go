package auth_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/clients/auth_client"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator"
)

type HandlerUser struct {
	clientUser auth_client.IClientUser
	jwtParser  jwt.ITokenParser
	validator  *validator.Validate
	logger     *logging.Logger
}

func NewHandlerUser(
	clientUser auth_client.IClientUser,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	logger *logging.Logger) *HandlerUser {

	return &HandlerUser{
		clientUser: clientUser,
		jwtParser:  jwtParser,
		validator:  validator,
		logger:     logger,
	}

}

func (h *HandlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("creating user...")

	w.Header().Set("Content-Type", "application/json")

	var createdUser CreateUserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&createdUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
	}

	if err := h.validator.Struct(createdUser); err != nil {
		h.logger.Warn("Failed to validate user data %w", err)
	}

	userID, err := h.clientUser.CreateUser(
		r.Context(),
		createdUser.Name,
		createdUser.Surname,
		createdUser.Email,
		createdUser.Password,
		createdUser.RepeatPassword,
		createdUser.Age,
	)

	if err != nil {
		h.logger.Warn("Failed to make request and register user %w", err)
	}

	resp, err := json.Marshal(userID)
	if err != nil {
		h.logger.Error("Failed to marshal userID %w", err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
