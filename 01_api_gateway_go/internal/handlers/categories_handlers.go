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

type HandlerCategories struct {
	clientCategories notes_server_clients.IClientCategories
	jwtParser        jwt.ITokenParser
	validator        *validator.Validate
	logger           *logging.Logger
}

func NewHandlerCategories(
	clientCategories notes_server_clients.IClientCategories,
	jwtParser jwt.ITokenParser,
	validator *validator.Validate,
	log *logging.Logger,

) *HandlerCategories {
	return &HandlerCategories{
		clientCategories: clientCategories,
		jwtParser:        jwtParser,
		validator:        validator,
		logger:           log,
	}
}

func (h *HandlerCategories) CreateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryCreate models.CategoryCreateRequest

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryCreate); err != nil {
		h.logger.Error("Failed to unmarshal category creation data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	categoryID, err := h.clientCategories.CreateCategory(r.Context(), &categoryCreate)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categoryID)

}

func (h *HandlerCategories) GetCategoryByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryGetByID models.CategoryGetRequestByID

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryGetByID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	category, err := h.clientCategories.GetCategoryByID(r.Context(), &categoryGetByID)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(category)

}

func (h *HandlerCategories) GetCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var userID int

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	categories, err := h.clientCategories.GetCategories(r.Context(), userID)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categories)

}

func (h *HandlerCategories) RemoveCategoryByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryRemoveByID models.CategoryRemoveRequestByID

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryRemoveByID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
	}

	categoryRemovedID, err := h.clientCategories.RemoveCategoryByID(r.Context(), &categoryRemoveByID)
	if err != nil {
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categoryRemovedID)

}
