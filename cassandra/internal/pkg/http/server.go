package http

import "net/http"

func NewServer(handler http.Handler, address string) *http.Server {
	return &http.Server{
		Handler: handler,
		Addr:    address,
	}
}
