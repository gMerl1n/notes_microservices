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
	UserContextKey      userCtx = "UserUUID"
)

func (h *HandlerUser) AuthMiddleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		header := request.Header.Get(AuthorizationHeader)

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			h.logger.Warn("headerParts", "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			h.logger.Warn("token is empty")
			return
		}

		userUUID, err := h.tokenManager.Parse(headerParts[1])
		if err != nil {
			h.logger.Warn("ParseToken", err.Error())
			return
		}

		ctx := context.WithValue(request.Context(), UserContextKey, userUUID)
		request = request.WithContext(ctx)

		hf(response, request)

	}
}
