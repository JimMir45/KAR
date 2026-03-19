package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/JimMir45/KAR/server/internal/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(c *gin.Context) {
	response.OK(c, gin.H{
		"status": "ok",
	})
}
