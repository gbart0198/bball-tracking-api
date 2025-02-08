package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (h *Handler) HandleGetGoalCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.repo.ListGoalCategories()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *Handler) HandleGetGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("goalCategoryId")
	categories := h.repo.GetGoalCategory(categoryId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (h *Handler) HandleCreateGoalCategory(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("goalCategoryId")
	createdCategory := h.repo.CreateGoalCategory(category)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdCategory)
}

func (h *Handler) HandleUpdateGoalCategory(w http.ResponseWriter, r *http.Request) {
	var updateGoalCategoryParams db.UpdateGoalCategoryParams

	err := json.NewDecoder(r.Body).Decode(&updateGoalCategoryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.repo.UpdateGoalCategory(updateGoalCategoryParams)
}

func (h *Handler) HandleDeleteGoalCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("categoryId")
	h.repo.DeleteGoalCategory(categoryId)
	w.WriteHeader(http.StatusNoContent)
}
