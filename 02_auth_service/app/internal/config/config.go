package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	BindAddr        string `yaml:"bind_ip"`
	Port            string `yaml:"port"`
	JWTSecret       string `yaml:"jwt_secret"`
	Debug           bool   `yaml:"is_debug"`
	Env             string `yaml:"env"`
	AccessTokenTTL  int    `yaml:"access_tokenTTL"`
	RefreshTokenTTL int    `yaml:"refresh_tokenTTL"`
}

func LoadConfig(p string) *Config {
	// path := fetchConfigPath()
	// if path == "" {
	// 	panic("config path is empty")
	// }

	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	panic("config file does not exist " + path)
	// }

	var cfg Config

	if err := cleanenv.ReadConfig(p, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath gets config path from command line flag or env variabl
// Priority: flag > env > default
// Default value is empty string
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
