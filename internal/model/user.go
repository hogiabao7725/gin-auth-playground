package model

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
	Role     string
}

var (
	ErrEmailTaken = errors.New("email already taken")
)

type UserRepository interface {
	Create(ctx context.Context, params CreateUserParams) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
}
