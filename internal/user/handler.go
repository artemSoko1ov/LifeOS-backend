package user

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data CreateUserRequest

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	newUser, err := h.service.CreateUser(r.Context(), data)
	if err != nil {
		switch {
		case errors.Is(err, ErrEmptyName),
			errors.Is(err, ErrEmptyEmail),
			errors.Is(err, ErrInvalidEmail),
			errors.Is(err, ErrEmptyLogin),
			errors.Is(err, ErrEmptyPassword):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, ErrEmailAlreadyExists),
			errors.Is(err, ErrLoginAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		return
	}
}
