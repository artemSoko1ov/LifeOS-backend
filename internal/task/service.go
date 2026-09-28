package task

import (
	"github.com/google/uuid"
	"time"
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

func (s *Service) CreateTask(title string) Task {
	now := time.Now()

	newTask := Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.repository.Create(newTask)

	return newTask
}
