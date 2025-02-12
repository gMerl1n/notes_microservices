package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
)

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) GetNoteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userId, ok := r.Context().Value(UserContextKey).(int)
	if !ok {
		h.logger.Error("Failed to parse userID from token")
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to parse userID from token")
	}

	var noteGetByID models.NoteGetRequestByID

	noteGetByID.UserID = userId

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&noteGetByID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	noteID, err := h.clientNotes.GetNoteByID(r.Context(), &noteGetByID)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteID)

}

func (h *Handler) GetNotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var notesGet models.NotesGetRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&notesGet); err != nil {
		h.logger.Error("Failed to unmarshal user ID  data to get notes %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	notes, err := h.clientNotes.GetNotes(r.Context(), &notesGet)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notes)

}

func (h *Handler) RemoveNoteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var noteRemoveByID models.NoteRemoveRequestByID

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&noteRemoveByID); err != nil {
		h.logger.Error("Failed to unmarshal note and user ID to remove a note %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	noteRemovedID, err := h.clientNotes.RemoveNoteByID(r.Context(), &noteRemoveByID)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteRemovedID)

}

func (h *Handler) RemoveNotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var notesRemove models.NotesRemove

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&notesRemove); err != nil {
		h.logger.Error("Failed to unmarshal user ID  data to remove notes %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	notes, err := h.clientNotes.RemoveNotes(r.Context(), &notesRemove)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notes)

}
