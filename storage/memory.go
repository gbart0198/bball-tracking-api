package storage

import (
	"github.com/gbart0198/bball-tracker-api/types"
	"github.com/google/uuid"
)

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) GetUser(userID uuid.UUID) *types.User {
	return &types.User{
		UserID:    userID,
		FirstName: "Test",
		LastName:  "User",
		Email:     "testuser@gmail.com",
	}
}
