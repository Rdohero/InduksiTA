package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SalesReport(c *gin.Context) {
	var Sales struct {
		Date     string `json:"date"`
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&Sales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", Sales.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid date format",
		})
		return
	}

	formattedDate := parsedDate.Format("2006-01-02T15:04:05Z")

	dateTime, err := time.Parse(time.RFC3339, formattedDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error converting date string to time.Time object",
		})
		return
	}

	Report := models.SalesReports{
		Date:     dateTime,
		ItemName: Sales.Item,
		Quantity: Sales.Quantity,
	}

	create := initializers.DB.Create(&Report).Find(&Report)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Sales Report",
			"Data":   Report,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func GetSalesReport(c *gin.Context) {
	var salesReport []models.SalesReports

	initializers.DB.Find(&salesReport)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Sales Report",
		"Data":   salesReport,
	})
}
