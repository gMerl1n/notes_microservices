package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	//"github.com/cristalhq/jwt"
	//"github.com/iriskin77/notes_microservices/app/internal/auth"
	//"github.com/iriskin77/notes_microservices/app/pkg/jwt"
)

func (h *Handler) AuthMiddleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		header := request.Header.Get("Authorization")

		headerParts := strings.Split(header, " ")

		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			fmt.Println("headerParts", "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			fmt.Println("token is empty")
			return
		}

		userId, err := h.tokenManager.Parse(headerParts[1])
		if err != nil {
			fmt.Println("ParseToken", err.Error())
			return
		}

		//fmt.Println(AuthMiddleware, userId)

		ctx := context.WithValue(request.Context(), "userId", userId)
		request = request.WithContext(ctx)

		hf(response, request)

	}
}
