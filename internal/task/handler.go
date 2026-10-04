package task

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
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

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	task, err := h.service.GetTaskByID(r.Context(), r.PathValue("id"))

	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "failed to encode task", http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var data UpdateTaskRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.service.UpdateTask(
		r.Context(),
		r.PathValue("id"),
		data,
	)
	if err != nil {
		if errors.Is(err, ErrEmptyTitle) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if errors.Is(err, ErrNothingToUpdate) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if errors.Is(err, ErrTaskNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	deletedTask, err := h.service.DeleteTask(r.Context(), r.PathValue("id"))

	if err != nil {
		if errors.Is(err, ErrTaskNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(deletedTask)
}
