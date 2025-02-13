package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/models"
	"github.com/gMerl1n/notes_microservices/pkg/apperrors"
)

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {

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

func (h *Handler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryGetByID models.CategoryGetRequestByID

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	categoryGetByID.UserID = numUserID

	h.logger.Info(fmt.Sprintf("Received user ID %d", numUserID))

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryGetByID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
		return
	}

	h.logger.Info(fmt.Sprintf("Received category ID %d", categoryGetByID.CategoryID))

	category, err := h.clientCategories.GetCategoryByID(r.Context(), &categoryGetByID)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get category by id from context %s ", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(category)

}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoriesGetRequest models.CategoriesGetRequest

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	categoriesGetRequest.UserID = numUserID

	categories, err := h.clientCategories.GetCategories(r.Context(), &categoriesGetRequest)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get category by id %s ", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categories)

}

func (h *Handler) RemoveCategoryByID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categoryRemoveByID models.CategoryRemoveRequestByID

	numUserID, err := h.GetContextUserID(r)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to get user ID from context %s ", err))
		return
	}

	categoryRemoveByID.UserID = numUserID

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&categoryRemoveByID); err != nil {
		h.logger.Error("Failed to unmarshal note get by id data %w", err)
		apperrors.BadRequestError(w, "Something wrong", 500, "Failed to decode login user data")
		return
	}

	categoryRemovedID, err := h.clientCategories.RemoveCategoryByID(r.Context(), &categoryRemoveByID)
	if err != nil {
		h.logger.Warn(err)
		apperrors.BadRequestError(w, "Something wrong", 500, fmt.Sprintf("Failed to remove category by id %s ", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(categoryRemovedID)

}
