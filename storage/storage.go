package storage

import (
	"github.com/gbart0198/bball-tracker-api/db"
)

type Storage interface {
    // Users Routes
    GetUser(string) *db.User
    ListUsers() []db.User
    CreateUser(db.CreateUserParams) *db.User
    UpdateUser(db.UpdateUserParams)
    DeleteUser(string)

    // Drills Routes
    GetDrill(string) *db.Drill
    ListDrills() []db.Drill
    CreateDrill(db.CreateDrillParams) *db.Drill
    UpdateDrill(db.UpdateDrillParams)
    DeleteDrill(string)

    // Goals Routes
    GetGoal(string) *db.Goal
    ListGoals() []db.Goal
    CreateGoal(db.CreateGoalParams) *db.Goal
    UpdateGoal(db.UpdateGoalParams)
    DeleteGoal(string)

    // Player Performances Routes
    GetPlayerPerformance(string) *db.PlayerPerformance
    ListPlayerPerformances() []db.PlayerPerformance
    CreatePlayerPerformance(db.CreatePerformanceParams) *db.PlayerPerformance
    UpdatePlayerPerformance(db.UpdatePerformanceParams)
    DeletePlayerPerformance(string)
    GetPerformancesByPlayer(string) []db.GetPerformancesByPlayerRow
    GetPerformancesByDrill(string) []db.GetPerformancesByDrillRow
    GetPerformancesBySession(string) []db.GetPerformancesBySessionRow

    // Session Routes
    GetSession(string) *db.Session
    ListSessions() []db.Session
    CreateSession(db.CreateSessionParams) *db.Session
    UpdateSession(db.UpdateSessionParams)
    DeleteSession(string)
    GetSessionsByOwner(string) []db.Session
    GetSessionByPerformance(string) *db.GetSessionByPerformanceRow

    // Player Goals Routes
    GetPlayerGoal(string) *db.PlayerGoal
    ListPlayerGoals() []db.PlayerGoal
    CreatePlayerGoal(db.CreatePlayerGoalParams) *db.PlayerGoal
    UpdatePlayerGoal(db.UpdatePlayerGoalParams)
    DeletePlayerGoal(string)
    GetGoalsByPlayer(string) []db.GetGoalsByPlayerRow

    // Session Performances Routes
    GetSessionPerformance(string) *db.SessionPerformance
    ListSessionPerformances() []db.SessionPerformance
    CreateSessionPerformance(db.CreateSessionPerformanceParams) *db.SessionPerformance
    UpdateSessionPerformance(db.UpdateSessionPerformanceParams)
    DeleteSessionPerformance(string)

    // Goal Category Routes
    GetGoalCategory(string) *db.GoalCategory
    ListGoalCategories() []db.GoalCategory
    CreateGoalCategory(string) *db.GoalCategory
    UpdateGoalCategory(db.UpdateGoalCategoryParams)
    DeleteGoalCategory(string)

}
