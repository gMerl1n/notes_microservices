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

type ConfigServer struct {
	Host     string
	Port     string
	LogLevel string
}

type Config struct {
	Server     *ConfigServer
	AuthServer *ConfigAuthServer
}

func fetchConfig() error {

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
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
	}, nil
}
