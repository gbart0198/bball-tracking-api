package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

/*
* --------------------------------------------------------
* User Handlers
* --------------------------------------------------------
 */

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	user := s.repo.GetUser(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleListUsers(w http.ResponseWriter, r *http.Request) {
	users := s.repo.ListUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user db.CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	createdUser := s.repo.CreateUser(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func (s *Server) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	var user db.UpdateUserParams
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	s.repo.UpdateUser(user)
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	s.repo.DeleteUser(userID)
}

/*
* --------------------------------------------------------
* Drills Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	drill := s.repo.GetDrill(drillID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drill)
}

func (s *Server) handleListDrills(w http.ResponseWriter, r *http.Request) {
	drills := s.repo.ListDrills()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drills)
}

func (s *Server) handleCreateDrill(w http.ResponseWriter, r *http.Request) {
	var drill db.CreateDrillParams
	err := json.NewDecoder(r.Body).Decode(&drill)
	if err != nil {
		log.Fatal(err)
	}
	createdDrill := s.repo.CreateDrill(drill)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdDrill)
}

func (s *Server) handleUpdateDrill(w http.ResponseWriter, r *http.Request) {
	var drill db.UpdateDrillParams
	err := json.NewDecoder(r.Body).Decode(&drill)
	if err != nil {
		log.Fatal(err)
	}
	s.repo.UpdateDrill(drill)
}

func (s *Server) handleDeleteDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	s.repo.DeleteDrill(drillID)
}

/*
* --------------------------------------------------------
* Goals Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("goalId")
	goal := s.repo.GetGoal(goalID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goal)
}

func (s *Server) handleListGoals(w http.ResponseWriter, r *http.Request) {
	goals := s.repo.ListGoals()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goals)
}

func (s *Server) handleCreateGoal(w http.ResponseWriter, r *http.Request) {
	var goal db.CreateGoalParams
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		log.Fatal(err)
	}
	createdGoal := s.repo.CreateGoal(goal)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdGoal)
}

func (s *Server) handleUpdateGoal(w http.ResponseWriter, r *http.Request) {
	var goal db.UpdateGoalParams
	err := json.NewDecoder(r.Body).Decode(&goal)
	if err != nil {
		log.Fatal(err)
	}
	s.repo.UpdateGoal(goal)
}

func (s *Server) handleDeleteGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("goalId")
	s.repo.DeleteGoal(goalID)
}

/*
* --------------------------------------------------------
* Player Performances Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetPlayerPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	performance := s.repo.GetPlayerPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performance)
}

func (s *Server) handleListPlayerPerformances(w http.ResponseWriter, r *http.Request) {
	performances := s.repo.ListPlayerPerformances()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (s *Server) handleCreatePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	var performance db.CreatePerformanceParams
	err := json.NewDecoder(r.Body).Decode(&performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdPerformance := s.repo.CreatePlayerPerformance(performance)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPerformance)
}

func (s *Server) handleUpdatePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	var performance db.UpdatePerformanceParams
	err := json.NewDecoder(r.Body).Decode(&performance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.repo.UpdatePlayerPerformance(performance)
}

func (s *Server) handleDeletePlayerPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	s.repo.DeletePlayerPerformance(performanceID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetPerformancesByPlayer(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	performances := s.repo.GetPerformancesByPlayer(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (s *Server) handleGetPerformancesByDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	performances := s.repo.GetPerformancesByDrill(drillID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

func (s *Server) handleGetPerformancesBySession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	performances := s.repo.GetPerformancesBySession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(performances)
}

/*
* --------------------------------------------------------
* Session Handlers
* --------------------------------------------------------
 */
func (s *Server) handleGetSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	session := s.repo.GetSession(sessionID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (s *Server) handleListSessions(w http.ResponseWriter, r *http.Request) {
	sessions := s.repo.ListSessions()

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
	createdSession := s.repo.CreateSession(session)
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
	s.repo.UpdateSession(session)
}

func (s *Server) handleDeleteSession(w http.ResponseWriter, r *http.Request) {
	sessionID := r.PathValue("sessionId")
	s.repo.DeleteSession(sessionID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetSessionsByOwner(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	sessions := s.repo.GetSessionsByOwner(userID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

func (s *Server) handleGetSessionByPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("performanceId")
	session := s.repo.GetSessionByPerformance(performanceID)

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
	playerGoal := s.repo.GetPlayerGoal(goalID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playerGoal)
}

func (s *Server) handleListPlayerGoals(w http.ResponseWriter, r *http.Request) {
	playerGoals := s.repo.ListPlayerGoals()

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
	createdPlayerGoal := s.repo.CreatePlayerGoal(playerGoal)
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
	s.repo.UpdatePlayerGoal(playerGoal)
}

func (s *Server) handleDeletePlayerGoal(w http.ResponseWriter, r *http.Request) {
	goalID := r.PathValue("playerGoalId")
	s.repo.DeletePlayerGoal(goalID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetGoalsByPlayer(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userId")
	goals := s.repo.GetGoalsByPlayer(userID)

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
	sessionPerformance := s.repo.GetSessionPerformance(performanceID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessionPerformance)
}

func (s *Server) handleListSessionPerformances(w http.ResponseWriter, r *http.Request) {
	sessionPerformances := s.repo.ListSessionPerformances()

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
	createdSessionPerformance := s.repo.CreateSessionPerformance(sessionPerformance)
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
	s.repo.UpdateSessionPerformance(sessionPerformance)
}

func (s *Server) handleDeleteSessionPerformance(w http.ResponseWriter, r *http.Request) {
	performanceID := r.PathValue("sessionPerformanceId")
	s.repo.DeleteSessionPerformance(performanceID)
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) handleGetGoalCategories(w http.ResponseWriter, r *http.Request) {
	categories := s.repo.ListGoalCategories()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (s *Server) handleGetGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("goalCategoryId")
	categories := s.repo.GetGoalCategory(categoryId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (s *Server) handleCreateGoalCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("goalCategoryId")
	createdCategory := s.repo.CreateGoalCategory(category)
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

	s.repo.UpdateGoalCategory(updateGoalCategoryParams)
}

func (s *Server) handleDeleteGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	s.repo.DeleteGoalCategory(categoryId)
	w.WriteHeader(http.StatusNoContent)
}
