package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConfigPostgres struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

func NewPostgresConfig(host string, port string, user string, password string, db_name string, sslmode string) *ConfigPostgres {

	return &ConfigPostgres{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		NameDB:   db_name,
		SSLMode:  sslmode}
}

func NewPostgresDB(ctx context.Context, cfg *ConfigPostgres) (dbpool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.NameDB)
	dbpool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool, nil

}
