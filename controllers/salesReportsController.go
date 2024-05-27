package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Items struct {
	ID              int    `json:"id"`
	Item            string `json:"item"`
	Price           int    `json:"price"`
	Category        string `json:"category"`
	CategoryItemsId int    `json:"category_items_id"`
	Quantity        int    `json:"quantity"`
}

func SalesReport(c *gin.Context) {
	var Sales struct {
		Date string  `json:"date"`
		Item []Items `json:"item"`
	}

	if err := c.BindJSON(&Sales); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if Sales.Item != nil && len(Sales.Item) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Sales Report",
			"Data":   Sales,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Item Not Found",
			"Data":  Sales,
		})
	}

	//parsedDate, err := time.Parse("2006-01-02", Sales.Date)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"Error": "Invalid date format",
	//	})
	//	return
	//}
	//
	//formattedDate := parsedDate.Format("2006-01-02T15:04:05Z")
	//
	//dateTime, err := time.Parse(time.RFC3339, formattedDate)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"Error": "Error converting date string to time.Time object",
	//	})
	//	return
	//}
	//
	//Report := models.SalesReports{
	//	Date:     dateTime,
	//}
	//
	//create := initializers.DB.Create(&Report).Find(&Report)
	//
	//if create.Error == nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"Succes": "Succes Create Sales Report",
	//		"Data":   Report,
	//	})
	//} else {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"error": create.Error,
	//	})
	//}
}

func GetSalesReport(c *gin.Context) {
	var salesReport []models.SalesReports

	initializers.DB.Preload("SalesReportItems.CategoryMachine").Find(&salesReport)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Sales Report",
		"Data":   salesReport,
	})
}
