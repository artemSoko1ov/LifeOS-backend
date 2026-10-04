package task

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(ctx context.Context) ([]Task, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, completed, created_at, updated_at
		FROM tasks
		ORDER BY created_at
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *Repository) Create(ctx context.Context, task Task) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO tasks (
			id,
			title,
			completed,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`,
		task.ID,
		task.Title,
		task.Completed,
		task.CreatedAt,
		task.UpdatedAt,
	)

	return err
}

func (r *Repository) GetByID(ctx context.Context, id string) (Task, error) {
	var task Task

	err := r.db.QueryRow(ctx, `
		SELECT id, title, completed, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`, id).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *Repository) Update(ctx context.Context, id string, data UpdateTaskRequest, newUpdatedAt time.Time) (Task, error) {
	var task Task

	err := r.db.QueryRow(ctx, `
		UPDATE tasks
		SET
		title = COALESCE($2, title),
		completed = COALESCE($3, completed),
		updated_at = $4
		WHERE id = $1
		RETURNING id, title, completed, created_at, updated_at
	`, id, data.Title, data.Completed, newUpdatedAt).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (r *Repository) Delete(ctx context.Context, id string) (Task, error) {
	var task Task

	err := r.db.QueryRow(ctx, `
		DELETE FROM tasks
		WHERE id = $1
		RETURNING id, title, completed, created_at, updated_at
	`, id).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}
