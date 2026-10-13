package user

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyName     = errors.New("name is required")
	ErrEmptyEmail    = errors.New("email is required")
	ErrInvalidEmail  = errors.New("invalid email format")
	ErrEmptyLogin    = errors.New("login is required")
	ErrEmptyPassword = errors.New("password is required")

	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrLoginAlreadyExists = errors.New("login already exists")
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateUser(ctx context.Context, request CreateUserRequest) (User, error) {
	if strings.TrimSpace(request.Name) == "" {
		return User{}, ErrEmptyName
	}

	if strings.TrimSpace(request.Email) == "" {
		return User{}, ErrEmptyEmail
	}

	if _, err := mail.ParseAddress(request.Email); err != nil {
		return User{}, ErrInvalidEmail
	}

	if strings.TrimSpace(request.Login) == "" {
		return User{}, ErrEmptyLogin
	}

	if strings.TrimSpace(request.Password) == "" {
		return User{}, ErrEmptyPassword
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}

	request.Name = strings.TrimSpace(request.Name)
	request.Email = strings.TrimSpace(request.Email)
	request.Login = strings.TrimSpace(request.Login)

	newUser := User{
		Name:         request.Name,
		Email:        request.Email,
		Login:        request.Login,
		PasswordHash: string(passwordHash),
	}

	user, err := s.repository.Create(ctx, newUser)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			switch pgErr.ConstraintName {
			case "users_email_key":
				return User{}, ErrEmailAlreadyExists
			case "users_login_key":
				return User{}, ErrLoginAlreadyExists
			}
		}

		return User{}, err
	}

	return user, nil
}
