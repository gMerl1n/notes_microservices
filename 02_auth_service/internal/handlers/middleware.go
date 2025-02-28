package handlers

import (
	"context"
	"net/http"
	"strings"
)

type userCtx string

const (
	// for AuthMiddleware
	AuthorizationHeader string  = "Authorization"
	UserContextKey      userCtx = "UserID"
)

func (h *HandlerUser) AuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get(AuthorizationHeader)

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			h.logger.Warn("headerParts", "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			h.logger.Warn("token is empty")
			return
		}

		userID, err := h.tokenManager.Parse(headerParts[1])
		if err != nil {
			h.logger.Warn("ParseToken", err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, userID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)

	}
}
