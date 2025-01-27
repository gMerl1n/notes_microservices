package apperrors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppError struct {
	Err              string `json:"error"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             int    `json:"code,omitempty"`
}

func newAppError(message string, code int, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Sprintf("Error: %s", message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func BadRequestError(w http.ResponseWriter, message string, code int, developerMessage string) {
	errorMessage := newAppError(message, code, developerMessage)

	sendErrorResponse(w, errorMessage)
}

func sendErrorResponse(w http.ResponseWriter, message *AppError) {
	responseJSON, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
