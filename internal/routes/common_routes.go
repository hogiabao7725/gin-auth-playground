package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/internal/handler"
)

type Handlers struct {
	Health *handler.HealthHandler
}

func SetupRoutes(router *gin.Engine, h Handlers) {
	registerHealthRoutes(router, h.Health)
}
