package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/app/internal/auth/services"
	"github.com/gMerl1n/notes_microservices/app/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/app/pkg/logging"
)

type Handler struct {
	services     services.IServiceUser
	tokenManager jwt.TokenManager
	logger       *logging.Logger
}

func NewHandler(services services.IServiceUser, tokenManager jwt.TokenManager, logger *logging.Logger) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		logger:       logger,
	}
}

func (h *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ok, err := json.Marshal("ok")
	if err != nil {
		h.logger.Error("Something wrong with the server")
	}

	h.logger.Fatal("The server is working. It is ready to accept requests")
	w.Write(ok)

}
