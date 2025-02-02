package api

import (
	"net/http"

	"github.com/gbart0198/bball-tracker-api/storage"
	"github.com/rs/cors"
)

type Server struct {
	listenAddr string
	repo       storage.Storage
}

func NewServer(listenAddr string, repo storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		repo:       repo,
	}
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	// User Routes
	router.HandleFunc("GET /user/{userId}", s.handleGetUser)
	router.HandleFunc("GET /user", s.handleListUsers)
	router.HandleFunc("PUT /user", s.handleCreateUser)
	router.HandleFunc("POST /user", s.handleUpdateUser)
	router.HandleFunc("DELETE /user/{userId}", s.handleDeleteUser)

	// Drill Routes
	router.HandleFunc("GET /drill/{drillId}", s.handleGetDrill)
	router.HandleFunc("GET /drill", s.handleListDrills)
	router.HandleFunc("PUT /drill", s.handleCreateDrill)
	router.HandleFunc("POST /drill", s.handleUpdateDrill)
	router.HandleFunc("DELETE /drill/{drillId}", s.handleDeleteDrill)

	// Goal Routes
	router.HandleFunc("GET /goal/{goalId}", s.handleGetGoal)
	router.HandleFunc("GET /goal", s.handleListGoals)
	router.HandleFunc("PUT /goal", s.handleCreateGoal)
	router.HandleFunc("POST /goal", s.handleUpdateGoal)
	router.HandleFunc("DELETE /goal/{goalId}", s.handleDeleteGoal)

	// Player Performance Routes
	router.HandleFunc("GET /performance/{performanceId}", s.handleGetPlayerPerformance)
	router.HandleFunc("GET /performance", s.handleListPlayerPerformances)
	router.HandleFunc("PUT /performance", s.handleCreatePlayerPerformance)
	router.HandleFunc("POST /performance", s.handleUpdatePlayerPerformance)
	router.HandleFunc("DELETE /performance/{performanceId}", s.handleDeletePlayerPerformance)
	router.HandleFunc("GET /performance/player/{userId}", s.handleGetPerformancesByPlayer)
	router.HandleFunc("GET /performance/drill/{drillId}", s.handleGetPerformancesByDrill)
	router.HandleFunc("GET /performance/session/{sessionId}", s.handleGetPerformancesBySession)

	// Session Routes
	router.HandleFunc("GET /session/{sessionId}", s.handleGetSession)
	router.HandleFunc("GET /session", s.handleListSessions)
	router.HandleFunc("PUT /session", s.handleCreateSession)
	router.HandleFunc("POST /session", s.handleUpdateSession)
	router.HandleFunc("DELETE /session/{sessionId}", s.handleDeleteSession)
	router.HandleFunc("GET /session/owner/{userId}", s.handleGetSessionsByOwner)
	router.HandleFunc("GET /session/performance/{performanceId}", s.handleGetSessionByPerformance)

	// Player Goal Routes
	router.HandleFunc("GET /player-goal/{playerGoalId}", s.handleGetPlayerGoal)
	router.HandleFunc("GET /player-goal", s.handleListPlayerGoals)
	router.HandleFunc("PUT /player-goal", s.handleCreatePlayerGoal)
	router.HandleFunc("POST /player-goal", s.handleUpdatePlayerGoal)
	router.HandleFunc("DELETE /player-goal/{playerGoalId}", s.handleDeletePlayerGoal)
	router.HandleFunc("GET /player-goal/player/{userId}", s.handleGetGoalsByPlayer)

	// Session Performance Routes
	router.HandleFunc("GET /session-performance/{sessionPerformanceId}", s.handleGetSessionPerformance)
	router.HandleFunc("GET /session-performance", s.handleListSessionPerformances)
	router.HandleFunc("PUT /session-performance", s.handleCreateSessionPerformance)
	router.HandleFunc("POST /session-performance", s.handleUpdateSessionPerformance)
	router.HandleFunc("DELETE /session-performance/{sessionPerformanceId}", s.handleDeleteSessionPerformance)

	apiRoot := http.NewServeMux()

	apiRoot.Handle("/api/", http.StripPrefix("/api", router))

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		Debug:          true,
	}).Handler(apiRoot)

	return http.ListenAndServe(s.listenAddr, handler)
}
