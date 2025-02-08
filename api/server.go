package api

import (
	"net/http"

	"github.com/gbart0198/bball-tracker-api/api/handlers"
	"github.com/gbart0198/bball-tracker-api/storage"
	"github.com/rs/cors"
)

type Server struct {
	listenAddr string
	handler    *handlers.Handler
}

func NewServer(listenAddr string, repo storage.Storage) *Server {
	handler := handlers.NewHandler(repo)
	return &Server{
		listenAddr: listenAddr,
		handler:    handler,
	}
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	// User Routes
	router.HandleFunc("GET /user/{userId}", s.handler.HandleGetUser)
	router.HandleFunc("GET /user", s.handler.HandleListUsers)
	router.HandleFunc("PUT /user", s.handler.HandleCreateUser)
	router.HandleFunc("POST /user", s.handler.HandleUpdateUser)
	router.HandleFunc("DELETE /user/{userId}", s.handler.HandleDeleteUser)

	// Drill Routes
	router.HandleFunc("GET /drill/{drillId}", s.handler.HandleGetDrill)
	router.HandleFunc("GET /drill", s.handler.HandleListDrills)
	router.HandleFunc("PUT /drill", s.handler.HandleCreateDrill)
	router.HandleFunc("POST /drill", s.handler.HandleUpdateDrill)
	router.HandleFunc("DELETE /drill/{drillId}", s.handler.HandleDeleteDrill)

	// Goal Routes
	router.HandleFunc("GET /goal/{goalId}", s.handler.HandleGetGoal)
	router.HandleFunc("GET /goal", s.handler.HandleListGoals)
	router.HandleFunc("PUT /goal", s.handler.HandleCreateGoal)
	router.HandleFunc("POST /goal", s.handler.HandleUpdateGoal)
	router.HandleFunc("DELETE /goal/{goalId}", s.handler.HandleDeleteGoal)

	// Player Performance Routes
	router.HandleFunc("GET /performance/{performanceId}", s.handler.HandleGetPlayerPerformance)
	router.HandleFunc("GET /performance", s.handler.HandleListPlayerPerformances)
	router.HandleFunc("PUT /performance", s.handler.HandleCreatePlayerPerformance)
	router.HandleFunc("POST /performance", s.handler.HandleUpdatePlayerPerformance)
	router.HandleFunc("DELETE /performance/{performanceId}", s.handler.HandleDeletePlayerPerformance)
	router.HandleFunc("GET /performance/player/{userId}", s.handler.HandleGetPerformancesByPlayer)
	router.HandleFunc("GET /performance/drill/{drillId}", s.handler.HandleGetPerformancesByDrill)
	router.HandleFunc("GET /performance/session/{sessionId}", s.handler.HandleGetPerformancesBySession)

	// Session Routes
	router.HandleFunc("GET /session/{sessionId}", s.handler.HandleGetSession)
	router.HandleFunc("GET /session", s.handler.HandleListSessions)
	router.HandleFunc("PUT /session", s.handler.HandleCreateSession)
	router.HandleFunc("POST /session", s.handler.HandleUpdateSession)
	router.HandleFunc("DELETE /session/{sessionId}", s.handler.HandleDeleteSession)
	router.HandleFunc("GET /session/owner/{userId}", s.handler.HandleGetSessionsByOwner)
	router.HandleFunc("GET /session/performance/{performanceId}", s.handler.HandleGetSessionByPerformance)

	// Player Goal Routes
	router.HandleFunc("GET /player-goal/{playerGoalId}", s.handler.HandleGetPlayerGoal)
	router.HandleFunc("GET /player-goal", s.handler.HandleListPlayerGoals)
	router.HandleFunc("PUT /player-goal", s.handler.HandleCreatePlayerGoal)
	router.HandleFunc("POST /player-goal", s.handler.HandleUpdatePlayerGoal)
	router.HandleFunc("DELETE /player-goal/{playerGoalId}", s.handler.HandleDeletePlayerGoal)
	router.HandleFunc("GET /player-goal/player/{userId}", s.handler.HandleGetGoalsByPlayer)

	// Session Performance Routes
	router.HandleFunc("GET /session-performance/{sessionPerformanceId}", s.handler.HandleGetSessionPerformance)
	router.HandleFunc("GET /session-performance", s.handler.HandleListSessionPerformances)
	router.HandleFunc("PUT /session-performance", s.handler.HandleCreateSessionPerformance)
	router.HandleFunc("POST /session-performance", s.handler.HandleUpdateSessionPerformance)
	router.HandleFunc("DELETE /session-performance/{sessionPerformanceId}", s.handler.HandleDeleteSessionPerformance)

	// Goal Category Routes
	router.HandleFunc("GET /goal-category/{goalCategoryId}", s.handler.HandleGetGoalCategory)
	router.HandleFunc("GET /goal-category", s.handler.HandleGetGoalCategories)
	router.HandleFunc("PUT /goal-category/{goalCategoryId}", s.handler.HandleCreateGoalCategory)
	router.HandleFunc("POST /goal-category", s.handler.HandleUpdateGoalCategory)
	router.HandleFunc("DELETE /goal-category/{goalCategoryId}", s.handler.HandleDeleteGoalCategory)

	apiRoot := http.NewServeMux()

	apiRoot.Handle("/api/", http.StripPrefix("/api", router))

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		Debug:          true,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler(apiRoot)

	return http.ListenAndServe(s.listenAddr, handler)
}
