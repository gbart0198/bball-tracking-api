// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createDrill = `-- name: CreateDrill :one
INSERT INTO drills (
    drill_name, category, difficulty
) VALUES (
    $1, $2, $3
)
RETURNING drill_id, drill_name, category, difficulty
`

type CreateDrillParams struct {
	DrillName  string      `json:"drillName"`
	Category   pgtype.Text `json:"category"`
	Difficulty pgtype.Text `json:"difficulty"`
}

func (q *Queries) CreateDrill(ctx context.Context, arg CreateDrillParams) (Drill, error) {
	row := q.db.QueryRow(ctx, createDrill, arg.DrillName, arg.Category, arg.Difficulty)
	var i Drill
	err := row.Scan(
		&i.DrillID,
		&i.DrillName,
		&i.Category,
		&i.Difficulty,
	)
	return i, err
}

const createGoal = `-- name: CreateGoal :one
INSERT INTO goals (
    goal_name, goal_type
) VALUES (
    $1, $2
)
RETURNING goal_id, goal_type, goal_name
`

type CreateGoalParams struct {
	GoalName string `json:"goalName"`
	GoalType string `json:"goalType"`
}

func (q *Queries) CreateGoal(ctx context.Context, arg CreateGoalParams) (Goal, error) {
	row := q.db.QueryRow(ctx, createGoal, arg.GoalName, arg.GoalType)
	var i Goal
	err := row.Scan(&i.GoalID, &i.GoalType, &i.GoalName)
	return i, err
}

const createGoalCategory = `-- name: CreateGoalCategory :one
INSERT INTO goal_categories (
    category
) VALUES (
    $1
)
RETURNING goal_category_id, category
`

func (q *Queries) CreateGoalCategory(ctx context.Context, category string) (GoalCategory, error) {
	row := q.db.QueryRow(ctx, createGoalCategory, category)
	var i GoalCategory
	err := row.Scan(&i.GoalCategoryID, &i.Category)
	return i, err
}

const createPerformance = `-- name: CreatePerformance :one
INSERT INTO player_performances (
    player_id, drill_id, date, attempts, successful
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING player_performance_id, player_id, drill_id, date, attempts, successful
`

type CreatePerformanceParams struct {
	PlayerID   uuid.UUID   `json:"playerId"`
	DrillID    uuid.UUID   `json:"drillId"`
	Date       pgtype.Date `json:"date"`
	Attempts   pgtype.Int4 `json:"attempts"`
	Successful pgtype.Int4 `json:"successful"`
}

func (q *Queries) CreatePerformance(ctx context.Context, arg CreatePerformanceParams) (PlayerPerformance, error) {
	row := q.db.QueryRow(ctx, createPerformance,
		arg.PlayerID,
		arg.DrillID,
		arg.Date,
		arg.Attempts,
		arg.Successful,
	)
	var i PlayerPerformance
	err := row.Scan(
		&i.PlayerPerformanceID,
		&i.PlayerID,
		&i.DrillID,
		&i.Date,
		&i.Attempts,
		&i.Successful,
	)
	return i, err
}

const createPlayerGoal = `-- name: CreatePlayerGoal :one
INSERT INTO player_goals (
    player_id, drill_id, current_value, goal_value, goal_category_id, goal_name, goal_description, completed
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING player_goal_id, player_id, drill_id, current_value, goal_value, goal_category_id, goal_name, goal_description, completed
`

type CreatePlayerGoalParams struct {
	PlayerID        uuid.UUID   `json:"playerId"`
	DrillID         uuid.UUID   `json:"drillId"`
	CurrentValue    pgtype.Int4 `json:"currentValue"`
	GoalValue       int32       `json:"goalValue"`
	GoalCategoryID  uuid.UUID   `json:"goalCategoryId"`
	GoalName        string      `json:"goalName"`
	GoalDescription pgtype.Text `json:"goalDescription"`
	Completed       bool        `json:"completed"`
}

func (q *Queries) CreatePlayerGoal(ctx context.Context, arg CreatePlayerGoalParams) (PlayerGoal, error) {
	row := q.db.QueryRow(ctx, createPlayerGoal,
		arg.PlayerID,
		arg.DrillID,
		arg.CurrentValue,
		arg.GoalValue,
		arg.GoalCategoryID,
		arg.GoalName,
		arg.GoalDescription,
		arg.Completed,
	)
	var i PlayerGoal
	err := row.Scan(
		&i.PlayerGoalID,
		&i.PlayerID,
		&i.DrillID,
		&i.CurrentValue,
		&i.GoalValue,
		&i.GoalCategoryID,
		&i.GoalName,
		&i.GoalDescription,
		&i.Completed,
	)
	return i, err
}

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
    user_id, session_type, session_name, date, location
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING session_id, session_type, date, location, user_id, session_name
`

type CreateSessionParams struct {
	UserID      uuid.UUID   `json:"userId"`
	SessionType string      `json:"sessionType"`
	SessionName string      `json:"sessionName"`
	Date        pgtype.Date `json:"date"`
	Location    pgtype.Text `json:"location"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.UserID,
		arg.SessionType,
		arg.SessionName,
		arg.Date,
		arg.Location,
	)
	var i Session
	err := row.Scan(
		&i.SessionID,
		&i.SessionType,
		&i.Date,
		&i.Location,
		&i.UserID,
		&i.SessionName,
	)
	return i, err
}

const createSessionPerformance = `-- name: CreateSessionPerformance :one
INSERT INTO session_performances (
    session_id, player_performance_id
) VALUES (
    $1, $2
)
RETURNING session_performance_id, session_id, player_performance_id
`

type CreateSessionPerformanceParams struct {
	SessionID           uuid.UUID `json:"sessionId"`
	PlayerPerformanceID uuid.UUID `json:"playerPerformanceId"`
}

func (q *Queries) CreateSessionPerformance(ctx context.Context, arg CreateSessionPerformanceParams) (SessionPerformance, error) {
	row := q.db.QueryRow(ctx, createSessionPerformance, arg.SessionID, arg.PlayerPerformanceID)
	var i SessionPerformance
	err := row.Scan(&i.SessionPerformanceID, &i.SessionID, &i.PlayerPerformanceID)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, passhash, email, firstname, lastname
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING user_id, username, passhash, email, firstname, lastname
`

type CreateUserParams struct {
	Username  string      `json:"username"`
	Passhash  string      `json:"passhash"`
	Email     string      `json:"email"`
	Firstname pgtype.Text `json:"firstname"`
	Lastname  pgtype.Text `json:"lastname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Passhash,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Passhash,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
	)
	return i, err
}

const deleteDrill = `-- name: DeleteDrill :exec
DELETE FROM drills
WHERE drill_id = $1
`

func (q *Queries) DeleteDrill(ctx context.Context, drillID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteDrill, drillID)
	return err
}

const deleteGoal = `-- name: DeleteGoal :exec
DELETE FROM goals
WHERE goal_id = $1
`

func (q *Queries) DeleteGoal(ctx context.Context, goalID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteGoal, goalID)
	return err
}

const deleteGoalCategory = `-- name: DeleteGoalCategory :exec
DELETE FROM goal_categories
WHERE goal_category_id = $1
`

func (q *Queries) DeleteGoalCategory(ctx context.Context, goalCategoryID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteGoalCategory, goalCategoryID)
	return err
}

const deletePerformance = `-- name: DeletePerformance :exec
DELETE FROM player_performances
WHERE player_performance_id = $1
`

func (q *Queries) DeletePerformance(ctx context.Context, playerPerformanceID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePerformance, playerPerformanceID)
	return err
}

const deletePlayerGoal = `-- name: DeletePlayerGoal :exec
DELETE FROM player_goals
WHERE player_goal_id = $1
`

func (q *Queries) DeletePlayerGoal(ctx context.Context, playerGoalID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deletePlayerGoal, playerGoalID)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE session_id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSession, sessionID)
	return err
}

const deleteSessionPerformance = `-- name: DeleteSessionPerformance :exec
DELETE FROM session_performances
WHERE session_performance_id = $1
`

func (q *Queries) DeleteSessionPerformance(ctx context.Context, sessionPerformanceID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteSessionPerformance, sessionPerformanceID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const getDrill = `-- name: GetDrill :one
SELECT drill_id, drill_name, category, difficulty FROM drills
WHERE drill_id = $1 LIMIT 1
`

func (q *Queries) GetDrill(ctx context.Context, drillID uuid.UUID) (Drill, error) {
	row := q.db.QueryRow(ctx, getDrill, drillID)
	var i Drill
	err := row.Scan(
		&i.DrillID,
		&i.DrillName,
		&i.Category,
		&i.Difficulty,
	)
	return i, err
}

const getGoal = `-- name: GetGoal :one
SELECT goal_id, goal_type, goal_name from goals
WHERE goal_id = $1 LIMIT 1
`

func (q *Queries) GetGoal(ctx context.Context, goalID uuid.UUID) (Goal, error) {
	row := q.db.QueryRow(ctx, getGoal, goalID)
	var i Goal
	err := row.Scan(&i.GoalID, &i.GoalType, &i.GoalName)
	return i, err
}

const getGoalCategories = `-- name: GetGoalCategories :many
SELECT goal_category_id, category from goal_categories
ORDER BY category asc
`

func (q *Queries) GetGoalCategories(ctx context.Context) ([]GoalCategory, error) {
	rows, err := q.db.Query(ctx, getGoalCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoalCategory
	for rows.Next() {
		var i GoalCategory
		if err := rows.Scan(&i.GoalCategoryID, &i.Category); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGoalCategory = `-- name: GetGoalCategory :one
SELECT goal_category_id, category from goal_categories
WHERE goal_category_id = $1 LIMIT 1
`

func (q *Queries) GetGoalCategory(ctx context.Context, goalCategoryID uuid.UUID) (GoalCategory, error) {
	row := q.db.QueryRow(ctx, getGoalCategory, goalCategoryID)
	var i GoalCategory
	err := row.Scan(&i.GoalCategoryID, &i.Category)
	return i, err
}

const getGoalsByPlayer = `-- name: GetGoalsByPlayer :many
SELECT
    pg.drill_id,
    pg.player_goal_id,
    d.drill_name,
    pg.goal_name,
    pg.goal_description,
    pg.current_value,
    pg.goal_value,
    gc.category,
    gc.goal_category_id,
    pg.completed
FROM player_goals as pg
JOIN drills d on d.drill_id = pg.drill_id
JOIN goal_categories gc on gc.goal_category_id = pg.goal_category_id
WHERE pg.player_id = $1
`

type GetGoalsByPlayerRow struct {
	DrillID         uuid.UUID   `json:"drillId"`
	PlayerGoalID    uuid.UUID   `json:"playerGoalId"`
	DrillName       string      `json:"drillName"`
	GoalName        string      `json:"goalName"`
	GoalDescription pgtype.Text `json:"goalDescription"`
	CurrentValue    pgtype.Int4 `json:"currentValue"`
	GoalValue       int32       `json:"goalValue"`
	Category        string      `json:"category"`
	GoalCategoryID  uuid.UUID   `json:"goalCategoryId"`
	Completed       bool        `json:"completed"`
}

func (q *Queries) GetGoalsByPlayer(ctx context.Context, playerID uuid.UUID) ([]GetGoalsByPlayerRow, error) {
	rows, err := q.db.Query(ctx, getGoalsByPlayer, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGoalsByPlayerRow
	for rows.Next() {
		var i GetGoalsByPlayerRow
		if err := rows.Scan(
			&i.DrillID,
			&i.PlayerGoalID,
			&i.DrillName,
			&i.GoalName,
			&i.GoalDescription,
			&i.CurrentValue,
			&i.GoalValue,
			&i.Category,
			&i.GoalCategoryID,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGoalsByPlayerAndDrill = `-- name: GetGoalsByPlayerAndDrill :many
SELECT
    pg.player_goal_id, pg.player_id, pg.drill_id, pg.current_value, pg.goal_value, pg.goal_category_id, pg.goal_name, pg.goal_description, pg.completed,
    gc.category
FROM player_goals as pg
JOIN drills d on d.drill_id = pg.drill_id
JOIN goal_categories gc on gc.goal_category_id = pg.goal_category_id
WHERE pg.player_id = $1 AND pg.drill_id = $2
`

type GetGoalsByPlayerAndDrillParams struct {
	PlayerID uuid.UUID `json:"playerId"`
	DrillID  uuid.UUID `json:"drillId"`
}

type GetGoalsByPlayerAndDrillRow struct {
	PlayerGoalID    uuid.UUID   `json:"playerGoalId"`
	PlayerID        uuid.UUID   `json:"playerId"`
	DrillID         uuid.UUID   `json:"drillId"`
	CurrentValue    pgtype.Int4 `json:"currentValue"`
	GoalValue       int32       `json:"goalValue"`
	GoalCategoryID  uuid.UUID   `json:"goalCategoryId"`
	GoalName        string      `json:"goalName"`
	GoalDescription pgtype.Text `json:"goalDescription"`
	Completed       bool        `json:"completed"`
	Category        string      `json:"category"`
}

func (q *Queries) GetGoalsByPlayerAndDrill(ctx context.Context, arg GetGoalsByPlayerAndDrillParams) ([]GetGoalsByPlayerAndDrillRow, error) {
	rows, err := q.db.Query(ctx, getGoalsByPlayerAndDrill, arg.PlayerID, arg.DrillID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGoalsByPlayerAndDrillRow
	for rows.Next() {
		var i GetGoalsByPlayerAndDrillRow
		if err := rows.Scan(
			&i.PlayerGoalID,
			&i.PlayerID,
			&i.DrillID,
			&i.CurrentValue,
			&i.GoalValue,
			&i.GoalCategoryID,
			&i.GoalName,
			&i.GoalDescription,
			&i.Completed,
			&i.Category,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPerformance = `-- name: GetPerformance :one
SELECT player_performance_id, player_id, drill_id, date, attempts, successful from player_performances
WHERE player_performance_id = $1 LIMIT 1
`

func (q *Queries) GetPerformance(ctx context.Context, playerPerformanceID uuid.UUID) (PlayerPerformance, error) {
	row := q.db.QueryRow(ctx, getPerformance, playerPerformanceID)
	var i PlayerPerformance
	err := row.Scan(
		&i.PlayerPerformanceID,
		&i.PlayerID,
		&i.DrillID,
		&i.Date,
		&i.Attempts,
		&i.Successful,
	)
	return i, err
}

const getPerformancesByDrill = `-- name: GetPerformancesByDrill :many
SELECT
    p.date,
    d.drill_id,
    d.drill_name,
    attempts,
    successful
FROM player_performances as p
JOIN drills as d on d.drill_id = p.drill_id
WHERE d.drill_id = $1
ORDER BY date desc
`

type GetPerformancesByDrillRow struct {
	Date       pgtype.Date `json:"date"`
	DrillID    uuid.UUID   `json:"drillId"`
	DrillName  string      `json:"drillName"`
	Attempts   pgtype.Int4 `json:"attempts"`
	Successful pgtype.Int4 `json:"successful"`
}

func (q *Queries) GetPerformancesByDrill(ctx context.Context, drillID uuid.UUID) ([]GetPerformancesByDrillRow, error) {
	rows, err := q.db.Query(ctx, getPerformancesByDrill, drillID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPerformancesByDrillRow
	for rows.Next() {
		var i GetPerformancesByDrillRow
		if err := rows.Scan(
			&i.Date,
			&i.DrillID,
			&i.DrillName,
			&i.Attempts,
			&i.Successful,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPerformancesByPlayer = `-- name: GetPerformancesByPlayer :many
SELECT
    player_performance_id,
    u.username,
    d.drill_name,
    p.date,
    attempts,
    successful
FROM player_performances as p
JOIN users as u ON p.player_id = u.user_id
JOIN drills as d on d.drill_id = p.drill_id
WHERE player_id = $1
ORDER BY date desc
`

type GetPerformancesByPlayerRow struct {
	PlayerPerformanceID uuid.UUID   `json:"playerPerformanceId"`
	Username            string      `json:"username"`
	DrillName           string      `json:"drillName"`
	Date                pgtype.Date `json:"date"`
	Attempts            pgtype.Int4 `json:"attempts"`
	Successful          pgtype.Int4 `json:"successful"`
}

func (q *Queries) GetPerformancesByPlayer(ctx context.Context, playerID uuid.UUID) ([]GetPerformancesByPlayerRow, error) {
	rows, err := q.db.Query(ctx, getPerformancesByPlayer, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPerformancesByPlayerRow
	for rows.Next() {
		var i GetPerformancesByPlayerRow
		if err := rows.Scan(
			&i.PlayerPerformanceID,
			&i.Username,
			&i.DrillName,
			&i.Date,
			&i.Attempts,
			&i.Successful,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPerformancesBySession = `-- name: GetPerformancesBySession :many
SELECT
    p.player_id,
    p.date,
    d.drill_id,
    d.drill_name,
    attempts,
    successful
FROM player_performances as p
JOIN drills as d on d.drill_id = p.drill_id
JOIN session_performances as sp on sp.player_performance_id = p.player_performance_id
WHERE sp.session_id = $1
`

type GetPerformancesBySessionRow struct {
	PlayerID   uuid.UUID   `json:"playerId"`
	Date       pgtype.Date `json:"date"`
	DrillID    uuid.UUID   `json:"drillId"`
	DrillName  string      `json:"drillName"`
	Attempts   pgtype.Int4 `json:"attempts"`
	Successful pgtype.Int4 `json:"successful"`
}

func (q *Queries) GetPerformancesBySession(ctx context.Context, sessionID uuid.UUID) ([]GetPerformancesBySessionRow, error) {
	rows, err := q.db.Query(ctx, getPerformancesBySession, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPerformancesBySessionRow
	for rows.Next() {
		var i GetPerformancesBySessionRow
		if err := rows.Scan(
			&i.PlayerID,
			&i.Date,
			&i.DrillID,
			&i.DrillName,
			&i.Attempts,
			&i.Successful,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayerGoal = `-- name: GetPlayerGoal :one
SELECT player_goal_id, player_id, drill_id, current_value, goal_value, goal_category_id, goal_name, goal_description, completed from player_goals
WHERE player_goal_id = $1 LIMIT 1
`

func (q *Queries) GetPlayerGoal(ctx context.Context, playerGoalID uuid.UUID) (PlayerGoal, error) {
	row := q.db.QueryRow(ctx, getPlayerGoal, playerGoalID)
	var i PlayerGoal
	err := row.Scan(
		&i.PlayerGoalID,
		&i.PlayerID,
		&i.DrillID,
		&i.CurrentValue,
		&i.GoalValue,
		&i.GoalCategoryID,
		&i.GoalName,
		&i.GoalDescription,
		&i.Completed,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT session_id, session_type, date, location, user_id, session_name from sessions
WHERE session_id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, sessionID uuid.UUID) (Session, error) {
	row := q.db.QueryRow(ctx, getSession, sessionID)
	var i Session
	err := row.Scan(
		&i.SessionID,
		&i.SessionType,
		&i.Date,
		&i.Location,
		&i.UserID,
		&i.SessionName,
	)
	return i, err
}

const getSessionByPerformance = `-- name: GetSessionByPerformance :one
SELECT
    sessions.session_id,
    session_type,
    session_name,
    location,
    sessions.date,
    sessions.user_id
from sessions
JOIN session_performances as sp on sp.session_id = sessions.session_id
WHERE sp.player_performance_id = $1
`

type GetSessionByPerformanceRow struct {
	SessionID   uuid.UUID   `json:"sessionId"`
	SessionType string      `json:"sessionType"`
	SessionName string      `json:"sessionName"`
	Location    pgtype.Text `json:"location"`
	Date        pgtype.Date `json:"date"`
	UserID      uuid.UUID   `json:"userId"`
}

func (q *Queries) GetSessionByPerformance(ctx context.Context, playerPerformanceID uuid.UUID) (GetSessionByPerformanceRow, error) {
	row := q.db.QueryRow(ctx, getSessionByPerformance, playerPerformanceID)
	var i GetSessionByPerformanceRow
	err := row.Scan(
		&i.SessionID,
		&i.SessionType,
		&i.SessionName,
		&i.Location,
		&i.Date,
		&i.UserID,
	)
	return i, err
}

const getSessionPerformance = `-- name: GetSessionPerformance :one
SELECT session_performance_id, session_id, player_performance_id from session_performances
WHERE session_performance_id = $1 LIMIT 1
`

func (q *Queries) GetSessionPerformance(ctx context.Context, sessionPerformanceID uuid.UUID) (SessionPerformance, error) {
	row := q.db.QueryRow(ctx, getSessionPerformance, sessionPerformanceID)
	var i SessionPerformance
	err := row.Scan(&i.SessionPerformanceID, &i.SessionID, &i.PlayerPerformanceID)
	return i, err
}

const getSessionsByOwner = `-- name: GetSessionsByOwner :many
SELECT session_id, session_type, date, location, user_id, session_name from sessions
WHERE user_id = $1
`

func (q *Queries) GetSessionsByOwner(ctx context.Context, userID uuid.UUID) ([]Session, error) {
	rows, err := q.db.Query(ctx, getSessionsByOwner, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.SessionID,
			&i.SessionType,
			&i.Date,
			&i.Location,
			&i.UserID,
			&i.SessionName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT user_id, username, passhash, email, firstname, lastname from users
where user_id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Passhash,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
	)
	return i, err
}

const listDrills = `-- name: ListDrills :many
SELECT drill_id, drill_name, category, difficulty from drills
ORDER BY drill_name asc
`

func (q *Queries) ListDrills(ctx context.Context) ([]Drill, error) {
	rows, err := q.db.Query(ctx, listDrills)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Drill
	for rows.Next() {
		var i Drill
		if err := rows.Scan(
			&i.DrillID,
			&i.DrillName,
			&i.Category,
			&i.Difficulty,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listGoals = `-- name: ListGoals :many
SELECT goal_id, goal_type, goal_name from goals
ORDER BY goal_name asc
`

func (q *Queries) ListGoals(ctx context.Context) ([]Goal, error) {
	rows, err := q.db.Query(ctx, listGoals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Goal
	for rows.Next() {
		var i Goal
		if err := rows.Scan(&i.GoalID, &i.GoalType, &i.GoalName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPerformances = `-- name: ListPerformances :many
SELECT player_performance_id, player_id, drill_id, date, attempts, successful from player_performances
ORDER BY date desc
`

func (q *Queries) ListPerformances(ctx context.Context) ([]PlayerPerformance, error) {
	rows, err := q.db.Query(ctx, listPerformances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlayerPerformance
	for rows.Next() {
		var i PlayerPerformance
		if err := rows.Scan(
			&i.PlayerPerformanceID,
			&i.PlayerID,
			&i.DrillID,
			&i.Date,
			&i.Attempts,
			&i.Successful,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPlayerGoals = `-- name: ListPlayerGoals :many
SELECT player_goal_id, player_id, drill_id, current_value, goal_value, goal_category_id, goal_name, goal_description, completed from player_goals
ORDER BY player_goal_id
`

func (q *Queries) ListPlayerGoals(ctx context.Context) ([]PlayerGoal, error) {
	rows, err := q.db.Query(ctx, listPlayerGoals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PlayerGoal
	for rows.Next() {
		var i PlayerGoal
		if err := rows.Scan(
			&i.PlayerGoalID,
			&i.PlayerID,
			&i.DrillID,
			&i.CurrentValue,
			&i.GoalValue,
			&i.GoalCategoryID,
			&i.GoalName,
			&i.GoalDescription,
			&i.Completed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessionPerformances = `-- name: ListSessionPerformances :many
SELECT session_performance_id, session_id, player_performance_id from session_performances
ORDER BY session_id
`

func (q *Queries) ListSessionPerformances(ctx context.Context) ([]SessionPerformance, error) {
	rows, err := q.db.Query(ctx, listSessionPerformances)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SessionPerformance
	for rows.Next() {
		var i SessionPerformance
		if err := rows.Scan(&i.SessionPerformanceID, &i.SessionID, &i.PlayerPerformanceID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessions = `-- name: ListSessions :many
SELECT session_id, session_type, date, location, user_id, session_name from sessions
ORDER BY session_name asc
`

func (q *Queries) ListSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.Query(ctx, listSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.SessionID,
			&i.SessionType,
			&i.Date,
			&i.Location,
			&i.UserID,
			&i.SessionName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, username, passhash, email, firstname, lastname from users
ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.Passhash,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDrill = `-- name: UpdateDrill :exec
UPDATE drills
    SET drill_name = $2,
    category = $3,
    difficulty = $4
WHERE drill_id = $1
`

type UpdateDrillParams struct {
	DrillID    uuid.UUID   `json:"drillId"`
	DrillName  string      `json:"drillName"`
	Category   pgtype.Text `json:"category"`
	Difficulty pgtype.Text `json:"difficulty"`
}

func (q *Queries) UpdateDrill(ctx context.Context, arg UpdateDrillParams) error {
	_, err := q.db.Exec(ctx, updateDrill,
		arg.DrillID,
		arg.DrillName,
		arg.Category,
		arg.Difficulty,
	)
	return err
}

const updateGoal = `-- name: UpdateGoal :exec
UPDATE goals
    SET goal_name = $2,
    goal_type = $3
WHERE goal_id = $1
`

type UpdateGoalParams struct {
	GoalID   uuid.UUID `json:"goalId"`
	GoalName string    `json:"goalName"`
	GoalType string    `json:"goalType"`
}

func (q *Queries) UpdateGoal(ctx context.Context, arg UpdateGoalParams) error {
	_, err := q.db.Exec(ctx, updateGoal, arg.GoalID, arg.GoalName, arg.GoalType)
	return err
}

const updateGoalCategory = `-- name: UpdateGoalCategory :exec
UPDATE goal_categories
    SET category = $2
WHERE goal_category_id = $1
`

type UpdateGoalCategoryParams struct {
	GoalCategoryID uuid.UUID `json:"goalCategoryId"`
	Category       string    `json:"category"`
}

func (q *Queries) UpdateGoalCategory(ctx context.Context, arg UpdateGoalCategoryParams) error {
	_, err := q.db.Exec(ctx, updateGoalCategory, arg.GoalCategoryID, arg.Category)
	return err
}

const updatePerformance = `-- name: UpdatePerformance :exec
UPDATE player_performances
    SET player_id = $2,
    drill_id = $3,
    date = $4,
    attempts = $5,
    successful = $6
WHERE player_performance_id = $1
`

type UpdatePerformanceParams struct {
	PlayerPerformanceID uuid.UUID   `json:"playerPerformanceId"`
	PlayerID            uuid.UUID   `json:"playerId"`
	DrillID             uuid.UUID   `json:"drillId"`
	Date                pgtype.Date `json:"date"`
	Attempts            pgtype.Int4 `json:"attempts"`
	Successful          pgtype.Int4 `json:"successful"`
}

func (q *Queries) UpdatePerformance(ctx context.Context, arg UpdatePerformanceParams) error {
	_, err := q.db.Exec(ctx, updatePerformance,
		arg.PlayerPerformanceID,
		arg.PlayerID,
		arg.DrillID,
		arg.Date,
		arg.Attempts,
		arg.Successful,
	)
	return err
}

const updatePlayerGoal = `-- name: UpdatePlayerGoal :exec
UPDATE player_goals
    SET player_id = $2,
    drill_id = $3,
    current_value = $4,
    goal_value = $5,
    goal_category_id = $6,
    goal_name = $7,
    goal_description = $8,
    completed = $9
WHERE player_goal_id = $1
`

type UpdatePlayerGoalParams struct {
	PlayerGoalID    uuid.UUID   `json:"playerGoalId"`
	PlayerID        uuid.UUID   `json:"playerId"`
	DrillID         uuid.UUID   `json:"drillId"`
	CurrentValue    pgtype.Int4 `json:"currentValue"`
	GoalValue       int32       `json:"goalValue"`
	GoalCategoryID  uuid.UUID   `json:"goalCategoryId"`
	GoalName        string      `json:"goalName"`
	GoalDescription pgtype.Text `json:"goalDescription"`
	Completed       bool        `json:"completed"`
}

func (q *Queries) UpdatePlayerGoal(ctx context.Context, arg UpdatePlayerGoalParams) error {
	_, err := q.db.Exec(ctx, updatePlayerGoal,
		arg.PlayerGoalID,
		arg.PlayerID,
		arg.DrillID,
		arg.CurrentValue,
		arg.GoalValue,
		arg.GoalCategoryID,
		arg.GoalName,
		arg.GoalDescription,
		arg.Completed,
	)
	return err
}

const updateSession = `-- name: UpdateSession :exec
UPDATE sessions
    SET user_id = $2,
    session_type = $3,
    session_name = $4,
    date = $5,
    location = $6
WHERE session_id = $1
`

type UpdateSessionParams struct {
	SessionID   uuid.UUID   `json:"sessionId"`
	UserID      uuid.UUID   `json:"userId"`
	SessionType string      `json:"sessionType"`
	SessionName string      `json:"sessionName"`
	Date        pgtype.Date `json:"date"`
	Location    pgtype.Text `json:"location"`
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) error {
	_, err := q.db.Exec(ctx, updateSession,
		arg.SessionID,
		arg.UserID,
		arg.SessionType,
		arg.SessionName,
		arg.Date,
		arg.Location,
	)
	return err
}

const updateSessionPerformance = `-- name: UpdateSessionPerformance :exec
UPDATE session_performances
    SET session_id = $2,
    player_performance_id = $3
WHERE session_performance_id = $1
`

type UpdateSessionPerformanceParams struct {
	SessionPerformanceID uuid.UUID `json:"sessionPerformanceId"`
	SessionID            uuid.UUID `json:"sessionId"`
	PlayerPerformanceID  uuid.UUID `json:"playerPerformanceId"`
}

func (q *Queries) UpdateSessionPerformance(ctx context.Context, arg UpdateSessionPerformanceParams) error {
	_, err := q.db.Exec(ctx, updateSessionPerformance, arg.SessionPerformanceID, arg.SessionID, arg.PlayerPerformanceID)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
    SET username = $2,
    passhash = $3,
    email = $4,
    firstname = $5,
    lastname = $6
WHERE user_id = $1
`

type UpdateUserParams struct {
	UserID    uuid.UUID   `json:"userId"`
	Username  string      `json:"username"`
	Passhash  string      `json:"passhash"`
	Email     string      `json:"email"`
	Firstname pgtype.Text `json:"firstname"`
	Lastname  pgtype.Text `json:"lastname"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.UserID,
		arg.Username,
		arg.Passhash,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
	)
	return err
}
