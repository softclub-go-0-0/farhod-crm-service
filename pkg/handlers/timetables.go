package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"log"
	"net/http"
	"time"
)

func (h *handler) CreateTimetable(c *gin.Context) {
	var timetable models.Timetable
	err := c.ShouldBindJSON(&timetable)
	if err != nil {
		log.Println("create timetable:", err)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation error",
			"errors":  err.Error(),
		})
		return
	}

	if h.DB.Create(&timetable).Error != nil {
		log.Println("inserting timetable data to DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"timetable":    timetable,
		"default_time": time.Time{},
	})
	// json.MarshalJson
}
