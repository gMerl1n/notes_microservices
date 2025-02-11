package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/clients/notes_server_clients"
	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
)

type HandlerNotes struct {
	clientNotes notes_server_clients.IClientNotes
	jwtParser   jwt.ITokenParser
	validator   *validator.Validate
	logger      *logging.Logger
}

func NewHandlerNotes(
	clientNotes notes_server_clients.IClientNotes,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	log *logging.Logger,

) *HandlerNotes {
	return &HandlerNotes{
		clientNotes: clientNotes,
		jwtParser:   jwtParser,
		validator:   validator,
		logger:      log,
	}
}

func (h *HandlerNotes) CreateNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var noteCreate models.NoteCreateRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&noteCreate); err != nil {
		h.logger.Error("Failed to unmarshal note creation data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	noteID, err := h.clientNotes.CreateNote(r.Context(), &noteCreate)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteID)

}
