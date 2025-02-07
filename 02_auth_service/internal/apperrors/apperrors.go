package apperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             int    `json:"code,omitempty"`
}

func NewAppError(message string, code int, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func BadRequestError(w http.ResponseWriter, message string, code int, developerMessage string) {
	errorMessage := NewAppError(message, code, developerMessage)

	SendErrorResponse(w, errorMessage)
}

func SendErrorResponse(w http.ResponseWriter, message *AppError) {
	responseJSON, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}

	w.WriteHeader(message.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
