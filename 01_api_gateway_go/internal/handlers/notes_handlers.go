package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
)

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var noteCreate models.NoteCreateRequest

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	noteCreate.UserID = numUserID

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

	var noteGetByID models.NoteGetRequestByID

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	noteGetByID.UserID = numUserID

	h.logger.Info(fmt.Sprintf("Received user ID %d", numUserID))

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&noteGetByID); err != nil {
		h.logger.Error(fmt.Sprintf("Failed to unmarshal note get by id data %d", err))
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode note id from the request")
		return
	}

	h.logger.Info(fmt.Sprintf("Received note ID %d", noteGetByID.NoteID))

	noteID, err := h.clientNotes.GetNoteByID(r.Context(), &noteGetByID)
	if err != nil {
		h.logger.Error(fmt.Sprintf("Failed to get note by id %d ", err))
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to get note by id after request to the note service")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(noteID)

}

func (h *Handler) GetNotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var notesGet models.NotesGetRequest

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	notesGet.UserID = numUserID

	notes, err := h.clientNotes.GetNotes(r.Context(), &notesGet)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notes)

}

func (h *Handler) RemoveNoteByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var noteRemoveByID models.NoteRemoveRequestByID

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	noteRemoveByID.UserID = numUserID

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

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	notesRemove.UserID = numUserID

	notes, err := h.clientNotes.RemoveNotes(r.Context(), &notesRemove)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(notes)

}
