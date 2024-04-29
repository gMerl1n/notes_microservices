package server

import (
	"net/http"
)

func NewHttpServer(addr string) (*http.Server, error) {

	return &http.Server{
		Addr: addr,
	}, nil
}
