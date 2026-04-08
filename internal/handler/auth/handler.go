package auth

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hogiabao7725/go-ticket-engine/internal/apperror"
	"github.com/hogiabao7725/go-ticket-engine/internal/model"
	"github.com/hogiabao7725/go-ticket-engine/internal/response"
)

type Service interface {
	Create(ctx context.Context, arg model.CreateUserParams) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateRole(ctx context.Context, id uuid.UUID, role string) error
}

type AuthHandler struct {
	service Service
}

func NewAuthHandler(service Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	// 1. Bind JSON request (check valid json syntax)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, apperror.New(apperror.CodeInvalidInput, "Invalid JSON format"))
		return
	}

	// 2. Custom validation
	if errs := req.Validate(); len(errs) > 0 {
		appErr := apperror.New(apperror.CodeInvalidInput, "Invalid input data").WithDetails(errs)
		response.Error(c, appErr)
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
