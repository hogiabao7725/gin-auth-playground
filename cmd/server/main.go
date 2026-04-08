package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hogiabao7725/go-ticket-engine/internal/config"
	"github.com/hogiabao7725/go-ticket-engine/internal/database"
	"github.com/hogiabao7725/go-ticket-engine/internal/handler/auth"
	"github.com/hogiabao7725/go-ticket-engine/internal/handler/health"
	userRepo "github.com/hogiabao7725/go-ticket-engine/internal/repository/user"
	"github.com/hogiabao7725/go-ticket-engine/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// initialize logger
	initLogger(cfg.Server.Env)

	// database connection
	pgPool, err := database.NewPostgresPool(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize postgres")
	}
	defer pgPool.Close()
	log.Info().Msg("successfully connected to postgres")

	// redis connection
	rd, err := database.NewRedisClient(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize redis")
	}
	defer func() {
		if err := rd.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close redis client")
		}
	}()
	log.Info().Msg("successfully connected to redis")

	// 1. Setup Repositories
	userRepository := userRepo.NewUserRepository(pgPool)

	// 2. Setup Services
	userService := service.NewUserService(userRepository)

	// set up router and start server
	router := gin.New()
	router.Use(gin.Recovery())

	// core api group
	v1 := router.Group("/api/v1")

	// 3. Mount Handlers
	healthHandler := health.NewHealthHandler()
	healthHandler.RegisterRoutes(v1)

	authHandler := auth.NewAuthHandler(userService)
	authHandler.RegisterRoutes(v1)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Info().
		Str("addr", addr).
		Str("env", cfg.Server.Env).
		Msg("starting server")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server error")
	}
}

func initLogger(env string) {
	zerolog.TimeFieldFormat = time.RFC3339
	if strings.EqualFold(env, "development") {
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		return
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}
