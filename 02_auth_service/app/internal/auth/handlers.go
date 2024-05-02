package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iriskin77/notes_microservices/app/internal/apperrors"
	"github.com/iriskin77/notes_microservices/app/pkg/jwt"
	"github.com/iriskin77/notes_microservices/app/pkg/logging"
)

type Handler struct {
	services     Service
	tokenManager jwt.TokenManager
	logger       *logging.Logger
}

func NewHandler(services Service, tokenManager jwt.TokenManager, logger *logging.Logger) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		logger:       logger,
	}
}

func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/api/auth/register", h.CreateUser).Methods("Post")
	router.HandleFunc("/api/auth/login", h.Login).Methods("Get")
	router.HandleFunc("/api/auth/refreshtokens", h.AuthMiddleware(h.RefreshTokens)).Methods("Get")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	h.logger.Info("creating user...")

	w.Header().Set("Content-Type", "application/json")

	var crUser CreateUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&crUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500)
	}

	userUUID, err := h.services.CreateUser(r.Context(), crUser)
	if err != nil {
		apperrors.BadRequestError(w, err.Error(), 500)
	}

	resp, err := json.Marshal(userUUID)
	if err != nil {
		h.logger.Error("Failed to marshal userUUID %w", err)
		apperrors.BadRequestError(w, "Internal server error", 500)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var loginUser LoginUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500)
	}

	tokens, err := h.services.Login(r.Context(), loginUser)
	if err != nil {
		apperrors.BadRequestError(w, err.Error(), 500)
	}

	tokenBytes, err := json.Marshal(tokens)
	if err != nil {
		h.logger.Error("Failed to marshal tokens login user %w", err)
		apperrors.BadRequestError(w, "invalid json schema", 500)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func (h *Handler) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var token RefreshTokensInput

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		apperrors.BadRequestError(w, "Failed to unmarshal refresh token", 500)

	}

	newTokens, err := h.services.RefreshTokens(r.Context(), token.Token)
	if err != nil {
		apperrors.BadRequestError(w, "Failed to get new tokens", 500)
	}

	tokenBytes, err := json.Marshal(newTokens)
	if err != nil {
		apperrors.BadRequestError(w, "Failed to marshal new tokens", 500)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}
