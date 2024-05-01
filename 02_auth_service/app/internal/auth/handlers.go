package auth

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iriskin77/notes_microservices/app/pkg/jwt"
)

type Handler struct {
	services     Service
	tokenManager jwt.TokenManager
	logger       *slog.Logger
}

func NewHandler(services Service, tokenManager jwt.TokenManager, logger *slog.Logger) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		logger:       logger,
	}
}

func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("/api/auth/createuser", h.CreateUser).Methods("Post")
	router.HandleFunc("/api/auth/user", h.Login).Methods("Get")
	router.HandleFunc("/api/auth/refresh_tokens", h.RefreshTokens).Methods("Get")
	router.HandleFunc("/api/auth/delete_user", h.AuthMiddleware(h.DeleteUser)).Methods("Delete")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var crUser CreateUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&crUser); err != nil {
		fmt.Println(err.Error())
		return
	}

	userUUID, err := h.services.CreateUser(r.Context(), crUser)
	if err != nil {
		fmt.Println(err.Error())
	}
	//w.Header().Set("Location", fmt.Sprintf("%s/%s", usersURL, userUUID))
	w.WriteHeader(http.StatusCreated)

	resp, err := json.Marshal(userUUID)

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Write(resp)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Body)

	var loginUser LoginUserDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(loginUser.Email, loginUser.Password)

	tokens, err := h.services.Login(r.Context(), loginUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	tokenBytes, err := json.Marshal(tokens)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func (h *Handler) RefreshTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var token RefreshTokensInput

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		fmt.Println(err.Error())
		return
	}

	newTokens, err := h.services.RefreshTokens(r.Context(), token.Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	tokenBytes, err := json.Marshal(newTokens)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	UserUUID := r.Context().Value("UserUUID").(string)

	deletedUserUUID, err := h.services.DeleteUser(r.Context(), UserUUID)
	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := json.Marshal(deletedUserUUID)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
