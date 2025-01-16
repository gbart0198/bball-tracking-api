package api

import (
	"net/http"
    "encoding/json"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
    userID := r.PathValue("userId")
	user := s.store.GetUser(userID)

	json.NewEncoder(w).Encode(user)
}
