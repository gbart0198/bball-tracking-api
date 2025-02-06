package storage

import (
	"context"
	"log"
	"log/slog"

	"github.com/gbart0198/bball-tracker-api/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type PostgreSqlStorage struct {
	queries *db.Queries
	ctx     context.Context
	pool    *pgxpool.Pool
}

func NewPostgreSqlStorage(ctx context.Context, connStr string) *PostgreSqlStorage {
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	queries := db.New(pool)
	slog.Info("DB connected")
	return &PostgreSqlStorage{
		ctx:     ctx,
		queries: queries,
		pool:    pool,
	}
}

func (store *PostgreSqlStorage) Close() {
	store.pool.Close()
}

/*
* ------------------------------------------------------------------
* User Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetUser(userID string) *db.User {
	user, err := store.queries.GetUser(store.ctx, uuid.MustParse(userID))
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

func (store *PostgreSqlStorage) ListUsers() []db.User {
	users, err := store.queries.ListUsers(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func (store *PostgreSqlStorage) CreateUser(params db.CreateUserParams) *db.User {
	insertedUser, err := store.queries.CreateUser(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedUser
}

func (store *PostgreSqlStorage) UpdateUser(params db.UpdateUserParams) {
	err := store.queries.UpdateUser(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

}

func (store *PostgreSqlStorage) DeleteUser(userID string) {
	err := store.queries.DeleteUser(store.ctx, uuid.MustParse(userID))
	if err != nil {
		log.Fatal(err)
	}
}

/*
* ------------------------------------------------------------------
* Drills Routes
* ------------------------------------------------------------------
 */
func (store *PostgreSqlStorage) GetDrill(drillID string) *db.Drill {
	drill, err := store.queries.GetDrill(store.ctx, uuid.MustParse(drillID))
	if err != nil {
		log.Fatal(err)
	}

	return &drill
}

func (store *PostgreSqlStorage) ListDrills() []db.Drill {
	drills, err := store.queries.ListDrills(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return drills
}

func (store *PostgreSqlStorage) CreateDrill(params db.CreateDrillParams) *db.Drill {
	insertedDrill, err := store.queries.CreateDrill(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedDrill
}

func (store *PostgreSqlStorage) UpdateDrill(params db.UpdateDrillParams) {
	err := store.queries.UpdateDrill(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

}

func (store *PostgreSqlStorage) DeleteDrill(drillID string) {
	err := store.queries.DeleteDrill(store.ctx, uuid.MustParse(drillID))
	if err != nil {
		log.Fatal(err)
	}
}

/*
* ------------------------------------------------------------------
* Goals Routes
* ------------------------------------------------------------------
 */
func (store *PostgreSqlStorage) GetGoal(goalID string) *db.Goal {
	goal, err := store.queries.GetGoal(store.ctx, uuid.MustParse(goalID))
	if err != nil {
		log.Fatal(err)
	}

	return &goal
}

func (store *PostgreSqlStorage) ListGoals() []db.Goal {
	goals, err := store.queries.ListGoals(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return goals
}

func (store *PostgreSqlStorage) CreateGoal(params db.CreateGoalParams) *db.Goal {
	createdGoal, err := store.queries.CreateGoal(store.ctx, params)
	if err != nil {
		log.Fatal(createdGoal)
	}

	return &createdGoal
}

func (store *PostgreSqlStorage) UpdateGoal(params db.UpdateGoalParams) {
	err := store.queries.UpdateGoal(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeleteGoal(goalID string) {
	err := store.queries.DeleteGoal(store.ctx, uuid.MustParse(goalID))
	if err != nil {
		log.Fatal(err)
	}
}

/*
* ------------------------------------------------------------------
* Player Performances Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetPlayerPerformance(playerPerformanceID string) *db.PlayerPerformance {
	playerPerformance, err := store.queries.GetPerformance(store.ctx, uuid.MustParse(playerPerformanceID))
	if err != nil {
		log.Fatal(err)
	}

	return &playerPerformance
}

func (store *PostgreSqlStorage) ListPlayerPerformances() []db.PlayerPerformance {
	playerPerformances, err := store.queries.ListPerformances(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return playerPerformances
}

func (store *PostgreSqlStorage) CreatePlayerPerformance(params db.CreatePerformanceParams) *db.PlayerPerformance {
	insertedPlayerPerformance, err := store.queries.CreatePerformance(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedPlayerPerformance
}

func (store *PostgreSqlStorage) UpdatePlayerPerformance(params db.UpdatePerformanceParams) {
	err := store.queries.UpdatePerformance(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeletePlayerPerformance(playerPerformanceID string) {
	err := store.queries.DeletePerformance(store.ctx, uuid.MustParse(playerPerformanceID))
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) GetPerformancesByPlayer(userID string) []db.GetPerformancesByPlayerRow {
	performances, err := store.queries.GetPerformancesByPlayer(store.ctx, uuid.MustParse(userID))
	if err != nil {
		log.Fatal(err)
	}

	return performances
}

func (store *PostgreSqlStorage) GetPerformancesByDrill(drillID string) []db.GetPerformancesByDrillRow {
	performances, err := store.queries.GetPerformancesByDrill(store.ctx, uuid.MustParse(drillID))
	if err != nil {
		log.Fatal(err)
	}

	return performances
}

func (store *PostgreSqlStorage) GetPerformancesBySession(sessionID string) []db.GetPerformancesBySessionRow {
	performances, err := store.queries.GetPerformancesBySession(store.ctx, uuid.MustParse(sessionID))
	if err != nil {
		log.Fatal(err)
	}

	return performances
}

/*
* ------------------------------------------------------------------
* Sessions Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetSession(sessionID string) *db.Session {
	session, err := store.queries.GetSession(store.ctx, uuid.MustParse(sessionID))
	if err != nil {
		log.Fatal(err)
	}

	return &session
}

func (store *PostgreSqlStorage) ListSessions() []db.Session {
	sessions, err := store.queries.ListSessions(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

func (store *PostgreSqlStorage) CreateSession(params db.CreateSessionParams) *db.Session {
	insertedSession, err := store.queries.CreateSession(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedSession
}

func (store *PostgreSqlStorage) UpdateSession(params db.UpdateSessionParams) {
	err := store.queries.UpdateSession(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeleteSession(sessionID string) {
	err := store.queries.DeleteSession(store.ctx, uuid.MustParse(sessionID))
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) GetSessionsByOwner(userID string) []db.Session {
	sessions, err := store.queries.GetSessionsByOwner(store.ctx, uuid.MustParse(userID))
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

func (store *PostgreSqlStorage) GetSessionByPerformance(performanceID string) *db.GetSessionByPerformanceRow {
	session, err := store.queries.GetSessionByPerformance(store.ctx, uuid.MustParse(performanceID))
	if err != nil {
		log.Fatal(err)
	}

	return &session
}

/*
* ------------------------------------------------------------------
* Player Goals Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetPlayerGoal(playerGoalID string) *db.PlayerGoal {
	playerGoal, err := store.queries.GetPlayerGoal(store.ctx, uuid.MustParse(playerGoalID))
	if err != nil {
		log.Fatal(err)
	}

	return &playerGoal
}

func (store *PostgreSqlStorage) ListPlayerGoals() []db.PlayerGoal {
	playerGoals, err := store.queries.ListPlayerGoals(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return playerGoals
}

func (store *PostgreSqlStorage) CreatePlayerGoal(params db.CreatePlayerGoalParams) *db.PlayerGoal {
	insertedPlayerGoal, err := store.queries.CreatePlayerGoal(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedPlayerGoal
}

func (store *PostgreSqlStorage) UpdatePlayerGoal(params db.UpdatePlayerGoalParams) {
	err := store.queries.UpdatePlayerGoal(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeletePlayerGoal(playerGoalID string) {
	err := store.queries.DeletePlayerGoal(store.ctx, uuid.MustParse(playerGoalID))
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) GetGoalsByPlayer(userID string) []db.GetGoalsByPlayerRow {
	goals, err := store.queries.GetGoalsByPlayer(store.ctx, uuid.MustParse(userID))
	if err != nil {
		log.Fatal(err)
	}

	return goals
}

/*
* ------------------------------------------------------------------
* Session Performances Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetSessionPerformance(sessionPerformanceID string) *db.SessionPerformance {
	sessionPerformance, err := store.queries.GetSessionPerformance(store.ctx, uuid.MustParse(sessionPerformanceID))
	if err != nil {
		log.Fatal(err)
	}

	return &sessionPerformance
}

func (store *PostgreSqlStorage) ListSessionPerformances() []db.SessionPerformance {
	sessionPerformances, err := store.queries.ListSessionPerformances(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return sessionPerformances
}

func (store *PostgreSqlStorage) CreateSessionPerformance(params db.CreateSessionPerformanceParams) *db.SessionPerformance {
	insertedSessionPerformance, err := store.queries.CreateSessionPerformance(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return &insertedSessionPerformance
}

func (store *PostgreSqlStorage) UpdateSessionPerformance(params db.UpdateSessionPerformanceParams) {
	err := store.queries.UpdateSessionPerformance(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeleteSessionPerformance(sessionPerformanceID string) {
	err := store.queries.DeleteSessionPerformance(store.ctx, uuid.MustParse(sessionPerformanceID))
	if err != nil {
		log.Fatal(err)
	}
}

/*
* ------------------------------------------------------------------
* Goal Category Routes
* ------------------------------------------------------------------
 */

func (store *PostgreSqlStorage) GetGoalCategory(goalCategoryID string) *db.GoalCategory {
	goalCategory, err := store.queries.GetGoalCategory(store.ctx, uuid.MustParse(goalCategoryID))
	if err != nil {
		log.Fatal(err)
	}

	return &goalCategory
}

func (store *PostgreSqlStorage) ListGoalCategories() []db.GoalCategory {
	goalCategories, err := store.queries.GetGoalCategories(store.ctx)
	if err != nil {
		log.Fatal(err)
	}

	return goalCategories
}

func (store *PostgreSqlStorage) CreateGoalCategory(goalCategoryName string) *db.GoalCategory {
	createdGoalCategory, err := store.queries.CreateGoalCategory(store.ctx, goalCategoryName)
	if err != nil {
		log.Fatal(err)
	}

	return &createdGoalCategory
}

func (store *PostgreSqlStorage) UpdateGoalCategory(params db.UpdateGoalCategoryParams) {
	err := store.queries.UpdateGoalCategory(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) DeleteGoalCategory(goalCategoryID string) {
	err := store.queries.DeleteGoalCategory(store.ctx, uuid.MustParse(goalCategoryID))
	if err != nil {
		log.Fatal(err)
	}
}

func (store *PostgreSqlStorage) GetGoalsByPlayerAndDrill(params db.GetGoalsByPlayerAndDrillParams) []db.GetGoalsByPlayerAndDrillRow {
	goals, err := store.queries.GetGoalsByPlayerAndDrill(store.ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	return goals
}
