package user

import (
	"context"

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

func (r *Repository) Create(ctx context.Context, user User) (User, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO users (
			name,
			email,
			login,
			password_hash
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, email, login, password_hash, created_at, updated_at
	`,
		user.Name,
		user.Email,
		user.Login,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Login,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
