package auth

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
	UserContextKey      userCtx = "UserUUID"
)

func (h *Handler) AuthMiddleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		header := request.Header.Get(AuthorizationHeader)

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			fmt.Println("headerParts", "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			fmt.Println("token is empty")
			return
		}

		userUUID, err := h.tokenManager.Parse(headerParts[1])
		if err != nil {
			fmt.Println("ParseToken", err.Error())
			return
		}

		//fmt.Println(AuthMiddleware, userId)

		ctx := context.WithValue(request.Context(), UserContextKey, userUUID)
		request = request.WithContext(ctx)

		hf(response, request)

	}
}
