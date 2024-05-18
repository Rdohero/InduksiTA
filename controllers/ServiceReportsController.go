package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ServiceReport(c *gin.Context) {
	var Service struct {
		Date          time.Time
		Branch        string
		Person        string
		MachineNumber string
		MachineName   string
		Quantity      int
		Complaints    string
		Status        string
	}

	if err := c.BindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Report := models.ServiceReports{
		Date:          Service.Date,
		Branch:        Service.Branch,
		PersonName:    Service.Person,
		MachineNumber: Service.MachineNumber,
		MachineName:   Service.MachineName,
		Complaints:    Service.Complaints,
		Status:        Service.Status,
	}

	create := initializers.DB.Create(&Report).Find(&Report)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Service Report",
			"Data":   Report,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func GetServiceReport(c *gin.Context) {
	var serviceReport []models.ServiceReports

	initializers.DB.Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Service Report",
		"Data":   serviceReport,
	})
}
