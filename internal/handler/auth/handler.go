package auth

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hogiabao7725/go-ticket-engine/internal/model"
	"github.com/hogiabao7725/go-ticket-engine/pkg/apperror"
	"github.com/hogiabao7725/go-ticket-engine/pkg/response"
)

type Service interface {
	Create(ctx context.Context, arg model.CreateUserParams) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
}

type UserHandler struct {
	service Service
}

func NewUserHandler(service Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// 1. Bind JSON request - Check for errors
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.Wrap(
			err,
			apperror.ErrInvalidInput.Code,
			apperror.ErrInvalidInput.Message,
			apperror.ErrInvalidInput.StatusCode,
		))
		return
	}

	// 2. Map request to service params
	arg := model.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     "user",
	}

	// 3. Call service layer
	user, err := h.service.Create(c.Request.Context(), arg)
	if err != nil {
		if errors.Is(err, model.ErrEmailTaken) {
			response.Error(c, apperror.ErrEmailTaken)
			return
		}

		response.Error(c, err)
		return
	}

	response.Created(c, RegisterResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	})

}
