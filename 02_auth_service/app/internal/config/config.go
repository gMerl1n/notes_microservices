package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

type ConfigRedis struct {
	AddrRedis     string
	PasswordRedis string
	DBRedis       int
}

type ConfigServer struct {
	Host     string
	Port     string
	LogLevel string
}

type ConfigToken struct {
	JWTsecret       string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

type ConfigUser struct {
	UserRoleID   int
	AdminRoleID  int
	SuperAdminID int
}

func fetchConfig() error {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()

}

type Config struct {
	Server   *ConfigServer
	Postgres *ConfigPostgres
	Redis    *ConfigRedis
	Token    *ConfigToken
	User     *ConfigUser
}

func NewConfig() (*Config, error) {

	if err := fetchConfig(); err != nil {
		fmt.Printf("error initialization config %s", err.Error())
		return nil, err
	}

	dbr, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		return nil, err
	}

	return &Config{

		Server: &ConfigServer{
			Host:     viper.GetString("server.host"),
			Port:     viper.GetString("server.port"),
			LogLevel: viper.GetString("server.log_level"),
		},
		Token: &ConfigToken{
			JWTsecret:       viper.GetString("token.jwt_secret"),
			AccessTokenTTL:  viper.GetInt("token.access_token_TTL"),
			RefreshTokenTTL: viper.GetInt("token.refresh_token_TTL"),
		},

		Postgres: &ConfigPostgres{
			User:     os.Getenv("POSTGRES_USER"),
			NameDB:   os.Getenv("POSTGRES_DB"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			SSLMode:  os.Getenv("SSLMODE"),
		},
		Redis: &ConfigRedis{
			AddrRedis:     os.Getenv("REDIS_PORT"),
			PasswordRedis: os.Getenv("REDIS_PASSWORD"),
			DBRedis:       dbr,
		},
		User: &ConfigUser{
			UserRoleID:   1,
			AdminRoleID:  2,
			SuperAdminID: 3,
		},
	}, nil
}
