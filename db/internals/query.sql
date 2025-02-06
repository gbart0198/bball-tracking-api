-- name: GetUser :one
SELECT * from users
where user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * from users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
    username, passhash, email, firstname, lastname
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
    SET username = $2,
    passhash = $3,
    email = $4,
    firstname = $5,
    lastname = $6
WHERE user_id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: GetPerformance :one
SELECT * from player_performances
WHERE player_performance_id = $1 LIMIT 1;

-- name: ListPerformances :many
SELECT * from player_performances
ORDER BY date desc;

-- name: CreatePerformance :one
INSERT INTO player_performances (
    player_id, drill_id, date, attempts, successful
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdatePerformance :exec
UPDATE player_performances
    SET player_id = $2,
    drill_id = $3,
    date = $4,
    attempts = $5,
    successful = $6
WHERE player_performance_id = $1;

-- name: DeletePerformance :exec
DELETE FROM player_performances
WHERE player_performance_id = $1;

-- name: GetDrill :one
SELECT * FROM drills
WHERE drill_id = $1 LIMIT 1;

-- name: ListDrills :many
SELECT * from drills
ORDER BY drill_name asc;

-- name: CreateDrill :one
INSERT INTO drills (
    drill_name, category, difficulty
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateDrill :exec
UPDATE drills
    SET drill_name = $2,
    category = $3,
    difficulty = $4
WHERE drill_id = $1;

-- name: DeleteDrill :exec
DELETE FROM drills
WHERE drill_id = $1;

-- name: GetPerformancesByPlayer :many
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
ORDER BY date desc;


-- name: GetPerformancesByDrill :many
SELECT 
    p.date,
    d.drill_id,
    d.drill_name,
    attempts,
    successful
FROM player_performances as p
JOIN drills as d on d.drill_id = p.drill_id
WHERE d.drill_id = $1
ORDER BY date desc;

-- name: GetGoal :one
SELECT * from goals
WHERE goal_id = $1 LIMIT 1;

-- name: ListGoals :many
SELECT * from goals
ORDER BY goal_name asc;

-- name: CreateGoal :one
INSERT INTO goals (
    goal_name, goal_type
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateGoal :exec
UPDATE goals
    SET goal_name = $2,
    goal_type = $3
WHERE goal_id = $1;

-- name: DeleteGoal :exec
DELETE FROM goals
WHERE goal_id = $1;

-- name: GetPlayerGoal :one
SELECT * from player_goals
WHERE player_goal_id = $1 LIMIT 1;

-- name: ListPlayerGoals :many
SELECT * from player_goals
ORDER BY player_goal_id;

-- name: CreatePlayerGoal :one
INSERT INTO player_goals (
    player_id, drill_id, current_value, goal_value, goal_category_id, goal_name, goal_description, completed
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdatePlayerGoal :exec
UPDATE player_goals
    SET player_id = $2,
    drill_id = $3,
    current_value = $4,
    goal_value = $5,
    goal_category_id = $6,
    goal_name = $7,
    goal_description = $8,
    completed = $9
WHERE player_goal_id = $1;

-- name: DeletePlayerGoal :exec
DELETE FROM player_goals
WHERE player_goal_id = $1;

-- name: GetSession :one
SELECT * from sessions
WHERE session_id = $1 LIMIT 1;

-- name: ListSessions :many
SELECT * from sessions
ORDER BY session_name asc;

-- name: CreateSession :one
INSERT INTO sessions (
    user_id, session_type, session_name, date, location
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateSession :exec
UPDATE sessions
    SET user_id = $2,
    session_type = $3,
    session_name = $4,
    date = $5,
    location = $6
WHERE session_id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE session_id = $1;

-- name: GetSessionPerformance :one
SELECT * from session_performances
WHERE session_performance_id = $1 LIMIT 1;

-- name: ListSessionPerformances :many
SELECT * from session_performances
ORDER BY session_id;

-- name: CreateSessionPerformance :one
INSERT INTO session_performances (
    session_id, player_performance_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateSessionPerformance :exec
UPDATE session_performances
    SET session_id = $2,
    player_performance_id = $3
WHERE session_performance_id = $1;

-- name: DeleteSessionPerformance :exec
DELETE FROM session_performances
WHERE session_performance_id = $1;

-- name: GetSessionsByOwner :many
SELECT * from sessions
WHERE user_id = $1;

-- name: GetPerformancesBySession :many
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
WHERE sp.session_id = $1;

-- name: GetSessionByPerformance :one
SELECT 
    sessions.session_id,
    session_type,
    session_name,
    location,
    sessions.date,
    sessions.user_id
from sessions
JOIN session_performances as sp on sp.session_id = sessions.session_id
WHERE sp.player_performance_id = $1;

-- name: GetGoalsByPlayer :many
SELECT 
    pg.drill_id,
    d.drill_name,
    pg.goal_name,
    pg.goal_description,
    pg.current_value,
    pg.goal_value,
    gc.category,
    gc.goal_category_id
FROM player_goals as pg
JOIN drills d on d.drill_id = pg.drill_id
JOIN goal_categories gc on gc.goal_category_id = pg.goal_category_id
WHERE pg.player_id = $1;


-- name: GetGoalCategories :many
SELECT * from goal_categories
ORDER BY category asc;

-- name: GetGoalCategory :one
SELECT * from goal_categories
WHERE goal_category_id = $1 LIMIT 1;

-- name: CreateGoalCategory :one
INSERT INTO goal_categories (
    category
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateGoalCategory :exec
UPDATE goal_categories
    SET category = $2
WHERE goal_category_id = $1;

-- name: DeleteGoalCategory :exec
DELETE FROM goal_categories
WHERE goal_category_id = $1;

-- name: GetGoalsByPlayerAndDrill :many
SELECT 
    pg.*,
    gc.category
FROM player_goals as pg
JOIN drills d on d.drill_id = pg.drill_id
JOIN goal_categories gc on gc.goal_category_id = pg.goal_category_id
WHERE pg.player_id = $1 AND pg.drill_id = $2;

