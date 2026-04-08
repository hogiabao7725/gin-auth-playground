package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hogiabao7725/go-ticket-engine/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
)

type userService struct {
	userRepo model.UserRepository
}

func New(userRepo model.UserRepository) model.UserRepository {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, arg model.CreateUserParams) (*model.User, error) {
	user, err := s.userRepo.Create(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
			return nil, model.ErrEmailTaken
		}

		return nil, fmt.Errorf("service: failed to create user: %w", err)
	}

	return user, nil
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (s *userService) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return nil
}
