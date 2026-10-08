package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
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
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return User{}, err
	}

	newUser := User{
		Name:         request.Name,
		Email:        request.Email,
		Login:        request.Login,
		PasswordHash: string(passwordHash),
	}

	user, err := s.repository.Create(ctx, newUser)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
