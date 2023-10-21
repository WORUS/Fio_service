package handler

import (
	"fio/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	lists := router.Group("/clients")
	{
		lists.POST("/", h.GetClients)
		lists.GET("/", h.GetClients)
		lists.GET("/:id", h.GetClients)
		lists.PUT("/:id", h.GetClients)
		lists.DELETE("/:id", h.GetClients)
	}

	return router
}
