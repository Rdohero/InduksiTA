package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(c *gin.Context) {
	var status []models.Status

	initializers.DB.Preload("ServiceReports").Find(&status)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Status",
		"Data":    status,
	})
}
