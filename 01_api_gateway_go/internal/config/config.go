package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigAuthServer struct {
	UrlCreateUser   string
	UrlLoginUser    string
	UrlRefreshToken string
}

type ConfigNotesServer struct {
	UrlCreateNote         string
	UrlGetNoteByID        string
	UrlGetNotes           string
	UrlRemoveNoteByID     string
	UrlRemoveNotes        string
	UrlCreateCategory     string
	UrlGetCategoryByID    string
	UrlGetCategories      string
	UrlRemoveCategoryByID string
}

type ConfigServer struct {
	Host     string
	Port     string
	LogLevel string
}

type ConfigToken struct {
	SigningKey string
}

type Config struct {
	Server      *ConfigServer
	AuthServer  *ConfigAuthServer
	NotesServer *ConfigNotesServer
	Token       *ConfigToken
}

func fetchConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()

}

func NewConfig() (*Config, error) {

	if err := fetchConfig(); err != nil {
		fmt.Printf("error initialization config %s", err.Error())
		return nil, err
	}

	return &Config{
		Server: &ConfigServer{
			Host:     viper.GetString("server.host"),
			Port:     viper.GetString("server.port"),
			LogLevel: viper.GetString("server.log_level"),
		},
		AuthServer: &ConfigAuthServer{
			UrlCreateUser:   viper.GetString("auth_server.create_user_url"),
			UrlLoginUser:    viper.GetString("auth_server.login_user_url"),
			UrlRefreshToken: viper.GetString("auth_server.refresh_token_url"),
		},
		NotesServer: &ConfigNotesServer{
			UrlCreateNote:         viper.GetString("notes_server.create_note_url"),
			UrlGetNoteByID:        viper.GetString("notes_server.get_note_by_id"),
			UrlGetNotes:           viper.GetString("notes_server.get_notes"),
			UrlRemoveNoteByID:     viper.GetString("notes_server.remove_note_by_id"),
			UrlRemoveNotes:        viper.GetString("notes_server.remove_notes"),
			UrlCreateCategory:     viper.GetString("notes_server.create_category"),
			UrlGetCategoryByID:    viper.GetString("notes_server.get_category_by_id"),
			UrlGetCategories:      viper.GetString("notes_server.get_categories"),
			UrlRemoveCategoryByID: viper.GetString("notes_server.remove_category_by_id"),
		},
		Token: &ConfigToken{
			SigningKey: viper.GetString("token.jwt_secret"),
		},
	}, nil
}
