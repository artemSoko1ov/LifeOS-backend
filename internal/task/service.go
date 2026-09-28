package task

import (
	"github.com/google/uuid"
	"time"
)

type Service struct {
	storage *Storage
}

func NewService(storage *Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) GetTasks() []Task {
	return s.storage.Tasks
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

	s.storage.Tasks = append(s.storage.Tasks, newTask)

	return newTask
}
