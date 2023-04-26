package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"log"
	"net/http"
)

// CreateGroup need to work
func (h *handler) CreateGroup(c *gin.Context) {
	var group models.Group
	err := c.ShouldBindJSON(&group)
	if err != nil {
		log.Println("create group:", err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})
		return
	}

	if h.DB.Create(&group).Error != nil {
		log.Println("inserting group data to DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, group)
}
