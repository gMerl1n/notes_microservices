package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type userCtx string

const (
	// for AuthMiddleware
	AuthorizationHeader string  = "Authorization"
	UserContextKey      userCtx = "UserID"
)

func (h *Handler) AuthMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		header := request.Header.Get(AuthorizationHeader)

		fmt.Println(header)

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

		ctx := context.WithValue(request.Context(), UserContextKey, userID)
		request = request.WithContext(ctx)

		handlerFunc(response, request)

	}
}
