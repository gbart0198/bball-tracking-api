package storage

import (
	"context"
	"log"

	"github.com/gbart0198/bball-tracker-api/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type PostgreSqlStorage struct {
    queries *db.Queries
    ctx context.Context
    conn *pgx.Conn
}

func NewPostgreSqlStorage(ctx context.Context) *PostgreSqlStorage {
	return &PostgreSqlStorage{
        ctx: ctx,
	}
}

func (store *PostgreSqlStorage) Connect(connStr string) {
    conn, err := pgx.Connect(store.ctx, connStr)
    if err != nil {
        log.Fatal(err)
    }
    store.conn = conn
    queries := db.New(store.conn)
    store.queries = queries
    log.Println("DB connected")
}

func (store *PostgreSqlStorage) Close() {
    store.conn.Close(store.ctx)
}

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

func (store *PostgreSqlStorage) CreateUser(user db.User) {
    params := db.CreateUserParams{
        Username: user.Username,
        Passhash: user.Passhash,
        Email: user.Email,
        Firstname: user.Firstname,
        Lastname: user.Lastname,
    }
    insertedUser, err := store.queries.CreateUser(store.ctx, params)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println(insertedUser)
}

func (store *PostgreSqlStorage) GetPerformancesByUser(userID string) []db.GetPerformancesByPlayerRow {
    performances, err := store.queries.GetPerformancesByPlayer(store.ctx, uuid.MustParse(userID))
    if err != nil {
        log.Fatal(err)
    }

    return performances
}
