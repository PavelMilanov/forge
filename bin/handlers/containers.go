package handlers

import (
	"net/http"

	"github.com/PavelMilanov/forge/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getContainers(c *gin.Context) {
	//project := c.Param("project")
	containers, err := core.GetContainers(h.cli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		//"project":    project,
		"containers": containers})
}
