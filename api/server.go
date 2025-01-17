package api

import (
	"github.com/gbart0198/bball-tracker-api/storage"
	"net/http"
)

type Server struct {
	listenAddr string
	store      storage.Storage
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *Server) Start() error {
    router := http.NewServeMux()

    router.HandleFunc("GET /user/{userId}", s.handleGetUser)
    router.HandleFunc("GET /user", s.handleListUsers)
    router.HandleFunc("PUT /user", s.handleCreateUser)
    router.HandleFunc("GET /user/performance/{userId}", s.handleGetPerformancesByUser)
	return http.ListenAndServe(s.listenAddr, router)
}

