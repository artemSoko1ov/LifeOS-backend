package task

import (
	"encoding/json"
	"net/http"
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

	tasks := h.service.GetTasks()

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

	newTask := h.service.CreateTask(data.Title)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}
