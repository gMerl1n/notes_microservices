package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("creating user...")

	w.Header().Set("Content-Type", "application/json")

	var createdUser models.CreateUserRequest
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&createdUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode new user data")
	}

	if err := h.validator.Struct(createdUser); err != nil {
		h.logger.Warn("Failed to validate user data %w", err)
	}

	userID, err := h.clientUser.CreateUser(r.Context(), &createdUser)

	if err != nil {
		h.logger.Warn("Failed to make request and register user %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to create a new user in DB")
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(userID)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var loginUser models.LoginUserRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}
	fmt.Println(loginUser)
	tokens, err := h.clientUser.LoginUser(r.Context(), &loginUser)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokens)

}

func (h *Handler) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var token models.RefreshTokensRequest

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
