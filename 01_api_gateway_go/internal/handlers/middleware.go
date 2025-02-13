package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type userCtx string

const (
	// for AuthMiddleware
	AuthorizationHeader string  = "Authorization"
	UserContextKey      userCtx = "UserID"
)

func (h *Handler) AuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get(AuthorizationHeader)

		h.logger.Info(fmt.Sprintf("Token from header %s", header))

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			h.logger.Warn("headerParts", "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			h.logger.Warn("token is empty")
			return
		}

		stripedToken := strings.TrimSpace(headerParts[1])

		userID, err := h.jwtParser.Parse(stripedToken)
		if err != nil {
			h.logger.Warn("ParseToken ", err.Error())
			return
		}

		r.Context()

		ctx := context.WithValue(r.Context(), UserContextKey, userID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)

	}
}

// Getting userID from request context
func (h *Handler) GetContextUserID(r *http.Request) (int, error) {

	userID, ok := r.Context().Value(UserContextKey).(string)
	if !ok {
		h.logger.Error("Failed to parse userID from token")
		return 0, fmt.Errorf("failed to parse userID from token")
	}

	numUserID, err := strconv.Atoi(userID)
	if err != nil {
		h.logger.Warn("Failed to conver string userID to int userID")
		return 0, fmt.Errorf("failed to convert string userID to int userID")
	}

	return numUserID, err

}
