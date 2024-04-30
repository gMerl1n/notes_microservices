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

	user, err := h.services.GetByEmailAndPassword(r.Context(), loginUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	token, err := h.tokenManager.NewJWT(user.UUID)
	if err != nil {
		fmt.Println(err.Error())
	}

	refreshToken, err := h.tokenManager.NewRefreshToken()
	if err != nil {
		fmt.Println(err.Error())
	}

	res := jwt.TokenResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}

	tokenBytes, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(tokenBytes)

}
