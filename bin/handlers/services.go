package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/PavelMilanov/forge/core"
	"github.com/gin-gonic/gin"
)

// getServices  возвращает всю информацию о доступных сущностях для заданного проекта.
func (h *Handler) getServices(c *gin.Context) {
	project := c.Param("project")
	dir := c.GetHeader("X-Project-Dir")
	file := c.GetHeader("X-Compose-File")
	if dir == "" || file == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны заголовки X-Project-Dir и X-Compose-File"})
		return
	}
	path := filepath.Join(dir, file)
	var compose core.DockerCompose
	if err := compose.Parse(path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"project":    project,
		"images":     compose.Images,
		"containers": compose.Containers,
	})
}
