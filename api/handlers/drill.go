package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gbart0198/bball-tracker-api/db"
)

func (h *Handler) HandleGetDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	drill := h.repo.GetDrill(drillID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drill)
}

func (h *Handler) HandleListDrills(w http.ResponseWriter, r *http.Request) {
	drills := h.repo.ListDrills()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drills)
}

func (h *Handler) HandleCreateDrill(w http.ResponseWriter, r *http.Request) {
	var drill db.CreateDrillParams
	err := json.NewDecoder(r.Body).Decode(&drill)
	if err != nil {
		log.Fatal(err)
	}
	createdDrill := h.repo.CreateDrill(drill)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdDrill)
}

func (h *Handler) HandleUpdateDrill(w http.ResponseWriter, r *http.Request) {
	var drill db.UpdateDrillParams
	err := json.NewDecoder(r.Body).Decode(&drill)
	if err != nil {
		log.Fatal(err)
	}
	h.repo.UpdateDrill(drill)
}

func (h *Handler) HandleDeleteDrill(w http.ResponseWriter, r *http.Request) {
	drillID := r.PathValue("drillId")
	h.repo.DeleteDrill(drillID)
}
