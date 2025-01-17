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

-- name: ListPerformance :many
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

