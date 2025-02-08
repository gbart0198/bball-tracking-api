package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

/*
* --------------------------------------------------------
* Session Performance Handlers
* --------------------------------------------------------
 */
func (h *Handler) HandleGetSessionPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("sessionPerformanceId")
	sessionPerformance := h.repo.GetSessionPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessionPerformance)
}

func (h *Handler) HandleListSessionPerformances(w http.ResponseWriter, r *http.Request) {
	sessionPerformances := h.repo.ListSessionPerformances()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessionPerformances)
}

func (h *Handler) HandleCreateSessionPerformance(w http.ResponseWriter, r *http.Request) {
	var sessionPerformance db.CreateSessionPerformanceParams
	err := json.NewDecoder(r.Body).Decode(&sessionPerformance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdSessionPerformance := h.repo.CreateSessionPerformance(sessionPerformance)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSessionPerformance)
}

func (h *Handler) HandleUpdateSessionPerformance(w http.ResponseWriter, r *http.Request) {
	var sessionPerformance db.UpdateSessionPerformanceParams
	err := json.NewDecoder(r.Body).Decode(&sessionPerformance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repo.UpdateSessionPerformance(sessionPerformance)
}

func (h *Handler) HandleDeleteSessionPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("sessionPerformanceId")
	h.repo.DeleteSessionPerformance(performanceID)
	w.WriteHeader(http.StatusNoContent)
}
