package auth_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/clients/auth_client"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
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
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode new user data")
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
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to create a new user in DB")
	}

	resp, err := json.Marshal(userID)
	if err != nil {
		h.logger.Error("Failed to marshal userID %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to marshal returned user ID")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (h *HandlerUser) LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var loginUser LoginUserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	tokens, err := h.clientUser.LoginUser(r.Context(), loginUser.Email, loginUser.Password)
	if err != nil {
	}

	tokenBytes, err := json.Marshal(tokens)
	if err != nil {
		h.logger.Error("Failed to marshal tokens login user %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to marshal user tokens")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func (h *HandlerUser) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var token RefreshTokensRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		h.logger.Error("Failed to decode tokens %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode user tokens")
	}

	newTokens, err := h.clientUser.RefreshTokens(r.Context(), token.RefreshToken)
	if err != nil {
		h.logger.Error("Failed to refresh tokens login user %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to refresh tokens")
	}

	tokenBytes, err := json.Marshal(newTokens)
	if err != nil {
		h.logger.Error("Failed to marshal tokens Refresh Tokens user %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to refresh tokens")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}
