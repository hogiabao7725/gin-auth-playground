package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/internal/modules/auth/domain"
	"github.com/hogiabao7725/go-ticket-engine/internal/modules/auth/usecase"
)

type AuthHandler struct {
	registerUC *usecase.RegisterUseCase
}

func NewAuthHandler(registerUC *usecase.RegisterUseCase) *AuthHandler {
	return &AuthHandler{
		registerUC: registerUC,
	}
}

func (h *AuthHandler) RegisterRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var reqHTTP registerRequestHTTP
	if err := c.ShouldBindJSON(&reqHTTP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// HTTP DTO -> usecase DTO
	req := usecase.RegisterRequest{
		Name:     reqHTTP.Name,
		Email:    reqHTTP.Email,
		Password: reqHTTP.Password,
	}

	resp, err := h.registerUC.Execute(c.Request.Context(), req)
	if err != nil {
		status, message := mapDomainErrorToHTTP(err)
		c.JSON(status, gin.H{"error": message})
		return
	}

	// usecase DTO -> HTTP DTO
	c.JSON(http.StatusCreated, registerResponseHTTP{
		ID:        resp.ID,
		Name:      resp.Name,
		Email:     resp.Email,
		Role:      resp.Role,
		CreatedAt: resp.CreatedAt,
	})
}

func mapDomainErrorToHTTP(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrUserNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return http.StatusConflict, err.Error()
	case errors.Is(err, domain.ErrInvalidCredentials):
		return http.StatusUnauthorized, err.Error()
	case errors.Is(err, domain.ErrWeakPassword),
		errors.Is(err, domain.ErrInvalidName),
		errors.Is(err, domain.ErrInvalidEmail):
		return http.StatusBadRequest, err.Error()
	default:
		return http.StatusInternalServerError, "internal server error"
	}
}
