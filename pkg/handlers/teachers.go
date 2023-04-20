package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/softclub-go-0-0/crm-service/pkg/models"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher
	result := h.DB.Find(&teachers)
	if result.Error != nil {
		log.Println("DB error - cannot find teachers:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, teachers)
}

func (h *handler) GetOneTeacher(c *gin.Context) {
	teacherID, err := strconv.Atoi(c.Param("teacherID"))
	if err != nil {
		log.Println("client error - bad teacherID param:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	var teacher models.Teacher
	result := h.DB.First(&teacher, teacherID)
	if result.Error != nil {
		log.Println("client error - cannot find teacher:", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, teacher)
}
