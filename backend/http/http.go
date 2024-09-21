package http

import (
	"encoding/json"
	"github.com/forthang/esketit/models"
	"github.com/go-chi/chi/v5"
	"log"
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

//var jwtSecret = []byte("superpupersecretkey")

func newServer(storage Storage) *Server {
	return &Server{storage: storage}
}

func (s *Server) MessagesHandler(w http.ResponseWriter, r *http.Request) {
	var jsonStruct struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&jsonStruct); err != nil {
		http.Error(w, "Error with parsing JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Println(jsonStruct.Message)

	jsonData, jsonErr := json.Marshal(jsonStruct)
	if jsonErr != nil {
		http.Error(w, "Error with json marshalling", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CreateAndRunServer(storage Storage, addr string) error {
	server := newServer(storage)

	r := chi.NewRouter()

	r.Post("/messages", server.MessagesHandler)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
