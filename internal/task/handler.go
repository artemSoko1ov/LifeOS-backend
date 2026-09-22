package task

import (
	"fmt"
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
	fmt.Fprintln(w, h.storage.tasks)
}
