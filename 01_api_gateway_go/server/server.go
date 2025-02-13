package server

import (
	"context"
	"net/http"

	"github.com/gMerl1n/notes_microservices/internal/clients"
	"github.com/gMerl1n/notes_microservices/internal/config"
	"github.com/gMerl1n/notes_microservices/internal/handlers"
	"github.com/gMerl1n/notes_microservices/pkg/client"
	"github.com/gMerl1n/notes_microservices/pkg/jwt"
	"github.com/gMerl1n/notes_microservices/pkg/logging"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func NewHttpServer(ctx context.Context, log *logging.Logger, conf *config.Config, validator *validator.Validate, jwtParser *jwt.TokenParser) (*http.Server, error) {

	baseClient := client.NewBaseClient(log)

	// Инициализация клиентов
	clients := clients.NewClient(baseClient, log, conf)

	// Инициализация ручек
	handlers := handlers.NewHandler(clients.UserClient, clients.NotesClient, clients.CategoriesClient, jwtParser, validator, log)

	router := mux.NewRouter()

	// auth handlers
	router.HandleFunc("/api_gateway/v1/create_user", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api_gateway/v1/login_user", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/api_gateway/v1/refresh_token", handlers.RefreshTokens).Methods("POST")

	// notices handlers
	router.HandleFunc("/api_gateway/v1/create_note", handlers.AuthMiddleware(handlers.CreateNote)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/get_note_by_id", handlers.AuthMiddleware(handlers.GetNoteByID)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/get_notes", handlers.AuthMiddleware(handlers.GetNotes)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/delete_note_by_id", handlers.AuthMiddleware(handlers.RemoveNoteByID)).Methods("DELETE")
	router.HandleFunc("/api_gateway/v1/delete_notes", handlers.AuthMiddleware(handlers.RemoveNotes)).Methods("DELETE")

	// categories handlers
	router.HandleFunc("/api_gateway/v1/create_category", handlers.AuthMiddleware(handlers.CreateCategory)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/get_category_by_id", handlers.AuthMiddleware(handlers.GetCategoryByID)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/get_categories", handlers.AuthMiddleware(handlers.GetCategories)).Methods("POST")
	router.HandleFunc("/api_gateway/v1/remove_category_by_id", handlers.AuthMiddleware(handlers.RemoveCategoryByID)).Methods("DELETE")

	return &http.Server{
		Addr:    conf.Server.Port,
		Handler: router,
	}, nil

}
