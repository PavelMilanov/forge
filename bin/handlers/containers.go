package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/PavelMilanov/forge/core"
	"github.com/gin-gonic/gin"
)

// getProjectContainers возвращает список контейнеров для проекта.
func (h *Handler) getContainers(c *gin.Context) {
	dir := c.GetHeader("X-Project-Dir")
	file := c.GetHeader("X-Compose-File")
	if dir == "" || file == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны заголовки X-Project-Dir и X-Compose-File"})
		return
	}
	path := filepath.Join(dir, file)
	containers, err := core.GetContainers(h.cli, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"containers": containers})
}

// getContainer возвращает информацию о контейнере.
func (h *Handler) getContainer(c *gin.Context) {
	id := c.Param("id")
	container, err := core.GetContainer(h.cli, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"container": container})
}

// startContainer запускает контейнер.
func (h *Handler) startContainer(c *gin.Context) {
	id := c.Param("id")
	err := core.StartContainer(h.cli, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Контейнер запущен"})
}

// restartContainer перезапускает контейнер.
func (h *Handler) restartContainer(c *gin.Context) {
	id := c.Param("id")
	err := core.RestartContainer(h.cli, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Контейнер перезапущен"})
}

// stopContainer останавливает контейнер.
func (h *Handler) stopContainer(c *gin.Context) {
	id := c.Param("id")
	err := core.StopContainer(h.cli, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Контейнер остановлен"})
}

// getLogsContainer возвращает логи контейнера.
func (h *Handler) getLogsContainer(c *gin.Context) {
	id := c.Param("id")
	logs, err := core.GetLogsContainer(h.cli, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"logs": string(logs)})
}
