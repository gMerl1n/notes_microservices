package auth

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	services Service
	logger   slog.Logger
}

func NewHandler(services Service, logger slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) RegisterHandlers(router *mux.Router) {
	router.HandleFunc("api/auth/createuser", h.CreateUser).Methods("Post")
	router.HandleFunc("api/auth/user/:uuid", h.GetUserByUUID).Methods("Get")
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

func (h *Handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

}
