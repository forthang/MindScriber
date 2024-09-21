package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Storage interface {
	Get(key string) (*storage.Task, error)
	Put(key string, value *storage.Task) error
	Post(key string, value *storage.Task) error
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
