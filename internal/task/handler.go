package task

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Handler struct {
	storage *Storage
}

type CreateTaskRequest struct {
	Title string `json:"title"`
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

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var data CreateTaskRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	title := data.Title
	now := time.Now()

	newTask := Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	h.storage.Tasks = append(h.storage.Tasks, newTask)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newTask)
}
