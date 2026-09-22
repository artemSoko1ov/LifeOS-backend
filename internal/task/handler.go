package task

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[]")
}
