package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    userID := r.PathValue("userId")
	user := s.store.GetUser(userID)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleListUsers(w http.ResponseWriter, r *http.Request) {
    users := s.store.ListUsers()

    json.NewEncoder(w).Encode(users)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user db.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        log.Fatal(err)
    }
    s.store.CreateUser(user)
}

func (s *Server) handleGetPerformancesByUser(w http.ResponseWriter, r *http.Request) {
    userID := r.PathValue("userId")
    performances := s.store.GetPerformancesByUser(userID)

    json.NewEncoder(w).Encode(performances)
}
