package handler

import (
	. "fio"
	"fio/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(client Client) *gin.Engine {
	router := gin.New()
	return router
}
