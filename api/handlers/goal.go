package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (h *Handler) HandleGetGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("goalId")
	goal := h.repo.GetGoal(goalID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goal)
}

func (h *Handler) HandleListGoals(w http.ResponseWriter, r *http.Request) {
	goals := h.repo.ListGoals()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

func (h *Handler) HandleCreateGoal(w http.ResponseWriter, r *http.Request) {
	var goal db.CreateGoalParams
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		log.Fatal(err)
	}
	createdGoal := h.repo.CreateGoal(goal)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdGoal)
}

func (h *Handler) HandleUpdateGoal(w http.ResponseWriter, r *http.Request) {
	var goal db.UpdateGoalParams
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		log.Fatal(err)
	}
	h.repo.UpdateGoal(goal)
}

func (h *Handler) HandleDeleteGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("goalId")
	h.repo.DeleteGoal(goalID)
}
