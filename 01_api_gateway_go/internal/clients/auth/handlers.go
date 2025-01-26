package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator"
)

type HandlerUser struct {
	authServer *config.ConfigAuthServer
	logger     *logging.Logger
}

func NewHandlerUser(log *logging.Logger, configAuthServer *config.ConfigAuthServer) *HandlerUser {
	return &HandlerUser{
		authServer: configAuthServer,
		validator * validator.Validate,
		logger: log,
	}
}

func (h *HandlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser CreateUserRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		h.logger.Error("Failed to unmarshal user data %w", err)
	}

	if err := h.validator.Struct(newUser); err != nil {
		h.logger.Warn("Failed to validate user data", err)
	}

}

func (h *HandlerUser) LoginUser(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerUser) RefreshTokens(w http.ResponseWriter, r *http.Request) {}
