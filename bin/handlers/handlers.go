package handlers

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

// Handler основная сущность взаимодействия с API.
type Handler struct {
	cli *client.Client
}

func NewHandler(cli *client.Client) *Handler {
	return &Handler{cli: cli}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1/")
	{
		services := v1.Group("/services/")
		{
			services.GET("/:project", h.getServices)
		}
		containers := v1.Group("/containers/")
		{
			containers.GET("/", h.getContainers)
		}
	}
	return router
}
