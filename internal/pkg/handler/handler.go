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
		lists.POST("/", h.CreateClient)
		lists.GET("/", h.GetClients)
		lists.GET("/:id", h.GetClients)
		lists.PUT("/:id", h.UpdateClientRecord)
		lists.DELETE("/:id", h.DeleteClientById)
	}

	return router
}
