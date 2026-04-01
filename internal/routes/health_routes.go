package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/internal/handler"
)

func registerHealthRoutes(router *gin.Engine, healthHandler *handler.HealthHandler) {
	router.GET(
		"/healthz",
		healthHandler.Healthz,
	)
}
