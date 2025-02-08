package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

/*
* --------------------------------------------------------
* Player Goals Handlers
* --------------------------------------------------------
 */
func (h *Handler) HandleGetPlayerGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("playerGoalId")
	playerGoal := h.repo.GetPlayerGoal(goalID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerGoal)
}

func (h *Handler) HandleListPlayerGoals(w http.ResponseWriter, r *http.Request) {
	playerGoals := h.repo.ListPlayerGoals()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerGoals)
}

func (h *Handler) HandleCreatePlayerGoal(w http.ResponseWriter, r *http.Request) {
	var playerGoal db.CreatePlayerGoalParams
	err := json.NewDecoder(r.Body).Decode(&playerGoal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPlayerGoal := h.repo.CreatePlayerGoal(playerGoal)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPlayerGoal)
}

func (s *Handler) HandleUpdatePlayerGoal(w http.ResponseWriter, r *http.Request) {
	var playerGoal db.UpdatePlayerGoalParams
	err := json.NewDecoder(r.Body).Decode(&playerGoal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.repo.UpdatePlayerGoal(playerGoal)
}

func (h *Handler) HandleDeletePlayerGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("playerGoalId")
	h.repo.DeletePlayerGoal(goalID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) HandleGetGoalsByPlayer(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	goals := h.repo.GetGoalsByPlayer(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}
