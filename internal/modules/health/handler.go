package health

import (
	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/internal/core/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/healthz", h.Healthz)
}

func (h *HealthHandler) Healthz(c *gin.Context) {
	data := gin.H{
		"status":  "ok",
		"message": "Ticket Engine is healthy",
	}
	response.OK(c, data)
}
