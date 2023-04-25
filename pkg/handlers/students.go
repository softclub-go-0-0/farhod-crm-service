package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"log"
	"net/http"
)

// CreateStudent need to work
func (h *handler) CreateStudent(c *gin.Context) {
	var student models.Student
	err := c.ShouldBindJSON(&student)
	if err != nil {
		log.Println("creating student:", err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "validation error",
			"err":     err.Error(),
		})
		return
	}

	if h.DB.Create(&student).Error != nil {
		log.Println("inserting student data to DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, student)
}
