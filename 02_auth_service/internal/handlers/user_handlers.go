package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/apperrors"
	"github.com/gMerl1n/notes_microservices/internal/services"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
)

type HandlerUser struct {
	services     services.IServiceUser
	tokenManager jwt.TokenManager
	logger       *logging.Logger
	validator    *validator.Validate
}

func NewHandlerUser(services services.IServiceUser, tokenManager jwt.TokenManager, logger *logging.Logger, validator *validator.Validate) *HandlerUser {
	return &HandlerUser{
		services:     services,
		tokenManager: tokenManager,
		logger:       logger,
		validator:    validator,
	}
}

func (h *HandlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("creating user...")

	w.Header().Set("Content-Type", "application/json")

	var createdUser CreateUserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&createdUser); err != nil {
		h.logger.Error("Failed to unmarshal user data: %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500, err.Error())
		return
	}

	if err := h.validator.Struct(createdUser); err != nil {
		h.logger.Warn("Failed to validate user data: %w", err)
		apperrors.BadRequestError(w, "Incorrect user data", 500, err.Error())
		return
	}

	userID, err := h.services.CreateUser(
		r.Context(),
		createdUser.Name,
		createdUser.Surname,
		createdUser.Email,
		createdUser.Password,
		createdUser.RepeatPassword,
		createdUser.Age,
	)

	if err != nil {
		h.logger.Warn("Failed to register new user %w", err)
		apperrors.BadRequestError(w, "failed to register user", 500, err.Error())
		return
	}

	resp, err := json.Marshal(userID)
	if err != nil {
		h.logger.Error("Failed to marshal userID %w", err)
		apperrors.BadRequestError(w, "Internal server error", 500, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (h *HandlerUser) LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var loginUser LoginUserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		h.logger.Error("Failed to unmarshal user data: %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500, err.Error())
		return

	}

	h.logger.Info("New user data: ", loginUser)

	tokens, err := h.services.Login(r.Context(), loginUser.Email, loginUser.Password)
	if err != nil {
		h.logger.Warn("Password does not match repeated password %w ", err)
		apperrors.BadRequestError(w, "Password does not match repeated password", 500, err.Error())
		return
	}

	tokenBytes, err := json.Marshal(tokens)
	if err != nil {
		h.logger.Error("Failed to marshal tokens login user %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func (h *HandlerUser) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var token RefreshTokensRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		h.logger.Warn("Failed to unmarshal refresh token %w", err)
		apperrors.BadRequestError(w, "Failed to unmarshal refresh token", 500, err.Error())
		return
	}

	newTokens, err := h.services.RefreshTokens(r.Context(), token.RefreshToken)
	if err != nil {
		h.logger.Warn("Failed to get new tokens %w", err)
		apperrors.BadRequestError(w, "Failed to get new tokens", 500, err.Error())
		return
	}

	tokenBytes, err := json.Marshal(newTokens)
	if err != nil {
		h.logger.Warn("Failed to get new tokens %w", err)
		apperrors.BadRequestError(w, "Failed to marshal new tokens", 500, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}
