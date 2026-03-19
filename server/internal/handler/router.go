package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/JimMir45/KAR/server/internal/config"
	"github.com/JimMir45/KAR/server/internal/middleware"
)

func SetupRouter(cfg *config.Config, logger *zap.Logger) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger(logger))
	r.Use(gin.Recovery())

	// Health check
	health := NewHealthHandler()
	r.GET("/api/health", health.Check)

	// API v1
	v1 := r.Group("/api/v1")

	// Public routes (no auth required)
	_ = v1.Group("/auth")
	{
		// TODO: POST /login — WeChat login
	}

	// Authenticated routes
	authed := v1.Group("")
	authed.Use(middleware.Auth(cfg.JWT.Secret))
	{
		// C-end routes
		_ = authed.Group("/activities")
		{
			// TODO: GET  /           — activity list
			// TODO: GET  /:id        — activity detail
			// TODO: POST /:id/register — register for activity
		}

		_ = authed.Group("/laps")
		{
			// TODO: GET /leaderboard — leaderboard by track
			// TODO: GET /my          — my lap records
			// TODO: GET /tracks      — track list
		}

		// B-end admin routes
		admin := authed.Group("/admin")
		admin.Use(middleware.RequireRole("ceo", "technician", "shareholder"))
		{
			_ = admin.Group("/activities")
			{
				// TODO: POST   /             — create activity
				// TODO: PUT    /:id          — update activity
				// TODO: POST   /:id/publish  — publish activity
				// TODO: GET    /:id/registrations — registration list
				// TODO: POST   /:id/laps     — batch input lap records
			}
		}
	}

	return r
}
