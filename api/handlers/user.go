package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	user := h.repo.GetUser(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) HandleListUsers(w http.ResponseWriter, r *http.Request) {
	users := h.repo.ListUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	createdUser := h.repo.CreateUser(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user db.UpdateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	h.repo.UpdateUser(user)
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	h.repo.DeleteUser(userID)
}
