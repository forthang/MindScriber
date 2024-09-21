package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Storage interface {
	Get(key string) (string, error)
	Put(key string, value string) error
	Post(key string, value string) error
	Delete(key string) error
}

type Server struct {
	storage Storage
}

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)

	r := chi.NewRouter()

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
