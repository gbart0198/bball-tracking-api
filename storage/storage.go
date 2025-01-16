package storage

import "github.com/gbart0198/bball-tracker-api/types"

type Storage interface {
    GetUser(string) *types.User
}
