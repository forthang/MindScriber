package http

import (
	"github.com/forthang/esketit/models"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Storage interface {
	CreateUser(user models.User) error
	GetUserByID(id uint) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type Server struct {
	storage Storage
}

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}

func (s *Server) TaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)

	r := chi.NewRouter()

	r.Get("/status", server.TaskStatusHandler)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
