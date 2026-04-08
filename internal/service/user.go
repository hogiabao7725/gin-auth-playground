package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/hogiabao7725/go-ticket-engine/internal/apperror"
	"github.com/hogiabao7725/go-ticket-engine/internal/model"
	"github.com/hogiabao7725/go-ticket-engine/pkg/hash"
	"github.com/jackc/pgx/v5/pgconn"
)

type userService struct {
	userRepo model.UserRepository
}

func NewUserService(userRepo model.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, arg model.CreateUserParams) (*model.User, error) {
	hashedPassword, err := hash.HashPassword(arg.Password)
	if err != nil {
		if errors.Is(err, hash.ErrPasswordEmpty) {
			return nil, apperror.New(apperror.CodeInvalidInput, "Password cannot be empty")
		}
		return nil, fmt.Errorf("service: failed to hash password: %w", err)
	}
	arg.Password = hashedPassword

	user, err := s.userRepo.Create(ctx, arg)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" && pgErr.ConstraintName == "users_email_key" {
			return nil, apperror.New(apperror.CodeEmailTaken, "this email is already exist.")
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
