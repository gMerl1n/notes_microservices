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

func BadRequestError(w http.ResponseWriter, message string, code int) {
	errorMessage := NewAppError(message, code, message)

	SendErrorResponse(w, errorMessage)
}

func SendErrorResponse(response http.ResponseWriter, message *AppError) {
	responseJSON, err := json.Marshal(message)
	if err != nil {
		http.Error(response, "Failed to marshal JSON", http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(responseJSON)
}
