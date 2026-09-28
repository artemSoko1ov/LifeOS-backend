package task

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTasks() []Task {
	return s.repository.Get()
}

func (s *Service) CreateTask(title string) (Task, error) {
	if strings.TrimSpace(title) == "" {
		return Task{}, errors.New("title is empty")
	}

	now := time.Now()

	newTask := Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repository.Create(newTask); err != nil {
		return Task{}, err
	}

	return newTask, nil
}
