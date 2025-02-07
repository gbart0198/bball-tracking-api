package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
	"github.com/gbart0198/bball-tracker-api/utils"
)

/*
* --------------------------------------------------------
* Session Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	session := s.handler.GetSession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (s *Server) handleListSessions(w http.ResponseWriter, r *http.Request) {
	sessions := s.handler.ListSessions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func (s *Server) handleCreateSession(w http.ResponseWriter, r *http.Request) {
	var session db.CreateSessionParams
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdSession := s.handler.CreateSession(session)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSession)
}

func (s *Server) handleUpdateSession(w http.ResponseWriter, r *http.Request) {
	var session db.UpdateSessionParams
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.handler.UpdateSession(session)
}

func (s *Server) handleDeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	s.handler.DeleteSession(sessionID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetSessionsByOwner(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	sessions := s.handler.GetSessionsByOwner(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func (s *Server) handleGetSessionByPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	session := s.handler.GetSessionByPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

/*
* --------------------------------------------------------
* Player Goals Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetPlayerGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("playerGoalId")
	playerGoal := s.handler.GetPlayerGoal(goalID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerGoal)
}

func (s *Server) handleListPlayerGoals(w http.ResponseWriter, r *http.Request) {
	playerGoals := s.handler.ListPlayerGoals()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerGoals)
}

func (s *Server) handleCreatePlayerGoal(w http.ResponseWriter, r *http.Request) {
	var playerGoal db.CreatePlayerGoalParams
	err := json.NewDecoder(r.Body).Decode(&playerGoal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPlayerGoal := s.handler.CreatePlayerGoal(playerGoal)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPlayerGoal)
}

func (s *Server) handleUpdatePlayerGoal(w http.ResponseWriter, r *http.Request) {
	var playerGoal db.UpdatePlayerGoalParams
	err := json.NewDecoder(r.Body).Decode(&playerGoal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.handler.UpdatePlayerGoal(playerGoal)
}

func (s *Server) handleDeletePlayerGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("playerGoalId")
	s.handler.DeletePlayerGoal(goalID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetGoalsByPlayer(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	goals := s.handler.GetGoalsByPlayer(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

/*
* --------------------------------------------------------
* Session Performance Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetSessionPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("sessionPerformanceId")
	sessionPerformance := s.handler.GetSessionPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessionPerformance)
}

func (s *Server) handleListSessionPerformances(w http.ResponseWriter, r *http.Request) {
	sessionPerformances := s.handler.ListSessionPerformances()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessionPerformances)
}

func (s *Server) handleCreateSessionPerformance(w http.ResponseWriter, r *http.Request) {
	var sessionPerformance db.CreateSessionPerformanceParams
	err := json.NewDecoder(r.Body).Decode(&sessionPerformance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdSessionPerformance := s.handler.CreateSessionPerformance(sessionPerformance)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdSessionPerformance)
}

func (s *Server) handleUpdateSessionPerformance(w http.ResponseWriter, r *http.Request) {
	var sessionPerformance db.UpdateSessionPerformanceParams
	err := json.NewDecoder(r.Body).Decode(&sessionPerformance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.handler.UpdateSessionPerformance(sessionPerformance)
}

func (s *Server) handleDeleteSessionPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("sessionPerformanceId")
	s.handler.DeleteSessionPerformance(performanceID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetGoalCategories(w http.ResponseWriter, r *http.Request) {
	categories := s.handler.ListGoalCategories()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (s *Server) handleGetGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("goalCategoryId")
	categories := s.handler.GetGoalCategory(categoryId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (s *Server) handleCreateGoalCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("goalCategoryId")
	createdCategory := s.handler.CreateGoalCategory(category)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCategory)
}

func (s *Server) handleUpdateGoalCategory(w http.ResponseWriter, r *http.Request) {
	var updateGoalCategoryParams db.UpdateGoalCategoryParams

	err := json.NewDecoder(r.Body).Decode(&updateGoalCategoryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.handler.UpdateGoalCategory(updateGoalCategoryParams)
}

func (s *Server) handleDeleteGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	s.handler.DeleteGoalCategory(categoryId)
	w.WriteHeader(http.StatusNoContent)
}
