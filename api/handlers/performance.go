package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
	"github.com/gbart0198/bball-tracker-api/utils"
)

/*
* --------------------------------------------------------
* Player Performances Handlers
* --------------------------------------------------------
 */
func (h *Handler) HandleGetPlayerPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	performance := h.repo.GetPlayerPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performance)
}

func (h *Handler) HandleListPlayerPerformances(w http.ResponseWriter, r *http.Request) {
	performances := h.repo.ListPlayerPerformances()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (h *Handler) HandleCreatePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	var performance db.CreatePerformanceParams
	err := json.NewDecoder(r.Body).Decode(&performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPerformance := h.repo.CreatePlayerPerformance(performance)
	utils.HandlePlayerGoalUpdates(createdPerformance, h.repo)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPerformance)
}

func (h *Handler) HandleUpdatePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	var performance db.UpdatePerformanceParams
	err := json.NewDecoder(r.Body).Decode(&performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repo.UpdatePlayerPerformance(performance)
}

func (h *Handler) HandleDeletePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	h.repo.DeletePlayerPerformance(performanceID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) HandleGetPerformancesByPlayer(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	performances := h.repo.GetPerformancesByPlayer(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (h *Handler) HandleGetPerformancesByDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	performances := h.repo.GetPerformancesByDrill(drillID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (h *Handler) HandleGetPerformancesBySession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	performances := h.repo.GetPerformancesBySession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}
