// Package handlers provides the HTTP handlers for the API.
//
// Each handler is responsible for handling a specific route and calling the appropriate function in the storage package.
package handlers

import (
	"github.com/gbart0198/bball-tracker-api/storage"
)

// A Handler is a struct that contains the storage repository and the associated functions for the API.
type Handler struct {
	repo storage.Storage
}

// NewHandler allocates and returns a new [Handler] with the given storage repository.
func NewHandler(repo storage.Storage) *Handler {
	return &Handler{repo}
}
