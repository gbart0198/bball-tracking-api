package storage

import (
	"github.com/gbart0198/bball-tracker-api/db"
)

type Storage interface {
    GetUser(string) *db.User
    ListUsers() []db.User
    CreateUser(db.User)
    GetPerformancesByUser(string) []db.GetPerformancesByPlayerRow
}
