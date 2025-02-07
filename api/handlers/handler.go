package handlers

import (
	"github.com/gbart0198/bball-tracker-api/storage"
)

type Handler struct {
	// The handler to call when the route is matched
	repo storage.Storage
}

func NewHandler(repo storage.Storage) *Handler {
	return &Handler{repo}
}
