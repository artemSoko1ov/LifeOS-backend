package task

import (
	"encoding/json"
	"net/http"
	"errors"
)

type Handler struct {
	service *Service
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks, err := h.service.GetTasks(r.Context())

	if err != nil {
		http.Error(w, "failed to get tasks", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
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

	newTask, err := h.service.CreateTask(r.Context(), data.Title)

if err != nil {
	if errors.Is(err, ErrEmptyTitle) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Error(w, "internal server error", http.StatusInternalServerError)
	return
}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}
