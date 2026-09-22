package task

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	storage *Storage
}

func NewHandler(storage *Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(h.storage.Tasks); err != nil {
		http.Error(w, "failed to encode tasks", http.StatusInternalServerError)
	}
}
