package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	sqlc "github.com/hogiabao7725/go-ticket-engine/internal/infra/sqlc"
	"github.com/hogiabao7725/go-ticket-engine/internal/model"
)

type adapter struct {
	q *sqlc.Queries
}

func NewUserRepository(db sqlc.DBTX) model.UserRepository {
	return &adapter{
		q: sqlc.New(db),
	}
}

// Implement the UserRepository interface defined in model/user.go

func (a *adapter) Create(ctx context.Context, arg model.CreateUserParams) (*model.User, error) {
	row, err := a.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
		Role:     sqlc.UserRole(arg.Role),
	})

	if err != nil {
		return nil, fmt.Errorf("repository: failed to create user: %w", err)
	}

	return &model.User{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		Role:      string(row.Role),
		CreatedAt: row.CreatedAt.Time,
	}, nil
}

func (a *adapter) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: implement and map from sqlc model to domain model.User
	return nil, nil
}

func (a *adapter) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return nil, nil
}

func (a *adapter) UpdateRole(ctx context.Context, id uuid.UUID, role string) error {
	return nil
}
