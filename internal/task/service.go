package task

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrEmptyTitle = errors.New("title is empty")
var ErrNothingToUpdate = errors.New("nothing to update.")
var ErrTaskNotFound = errors.New("task not found")

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

func (s *Service) GetTaskById(ctx context.Context, id string) (Task, error) {
	return s.repository.GetById(ctx, id)
}

func (s *Service) UpdateTask(ctx context.Context, id string, data UpdateTaskRequest) (Task, error) {
	if data.Title == nil && data.Completed == nil {
		return Task{}, ErrNothingToUpdate
	}

	if data.Title != nil {
		if strings.TrimSpace(*data.Title) == "" {
			return Task{}, ErrEmptyTitle
		}
	}

	newUpdatedAt := time.Now()

	updatedTask, err := s.repository.Update(ctx, id, data, newUpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return Task{}, ErrTaskNotFound
	}

	if err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}
