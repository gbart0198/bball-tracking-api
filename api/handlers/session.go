package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (h *Handler) HandleGetSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	session := h.repo.GetSession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (h *Handler) HandleListSessions(w http.ResponseWriter, r *http.Request) {
	sessions := h.repo.ListSessions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func (h *Handler) HandleCreateSession(w http.ResponseWriter, r *http.Request) {
	var session db.CreateSessionParams
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdSession := h.repo.CreateSession(session)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSession)
}

func (h *Handler) HandleUpdateSession(w http.ResponseWriter, r *http.Request) {
	var session db.UpdateSessionParams
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.repo.UpdateSession(session)
}

func (h *Handler) HandleDeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	h.repo.DeleteSession(sessionID)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) HandleGetSessionsByOwner(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	sessions := h.repo.GetSessionsByOwner(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func (h *Handler) HandleGetSessionByPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	session := h.repo.GetSessionByPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}
