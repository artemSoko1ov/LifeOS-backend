package task

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var ErrEmptyTitle = errors.New("title is empty")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetTasks(ctx context.Context) ([]Task, error) {
	return s.repository.Get(ctx)
}

func (s *Service) CreateTask(ctx context.Context, title string) (Task, error) {
	if strings.TrimSpace(title) == "" {
		return Task{}, ErrEmptyTitle
	}

	now := time.Now()

	newTask := Task{
		ID:        uuid.NewString(),
		Title:     title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repository.Create(ctx, newTask); err != nil {
		return Task{}, err
	}

	return newTask, nil
}
