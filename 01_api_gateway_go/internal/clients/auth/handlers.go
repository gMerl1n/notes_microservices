package auth

import (
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
)

type HandlerUser struct {
	authServer *config.ConfigAuthServer
	logger     *logging.Logger
}

func NewHandlerUser(log *logging.Logger, configAuthServer *config.ConfigAuthServer) *HandlerUser {
	return &HandlerUser{
		authServer: configAuthServer,
		logger:     log,
	}
}

func (h *HandlerUser) CreateUser() {}

func (h *HandlerUser) LoginUser() {}

func (h *HandlerUser) RefreshTokens() {}
