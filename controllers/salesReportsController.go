package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
			Date: dateTime,
		}

		create := initializers.DB.Create(&Report).Find(&Report)

		if create.Error == nil {
			for _, item := range Sales.Item {
				if item.Category == "mesin" {
					var storeItem models.StoreItems
					if store := initializers.DB.First(&storeItem, item.ID).Error; store != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"Error": "Store item not found",
							"Item":  item,
						})
						return
					}

					if storeItem.Quantity < item.Quantity {
						c.JSON(http.StatusBadRequest, gin.H{
							"Error": "Insufficient stock",
							"Item":  item,
						})
						return
					}
					storeItem.Quantity -= item.Quantity
					if quantityUpdate := initializers.DB.Save(&storeItem).Error; quantityUpdate != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"Error": "Failed to update store item quantity",
							"Item":  item,
						})
						return
					}
				} else if item.Category == "spare_part" {
					var sparePart models.SparePart
					if spare := initializers.DB.First(&sparePart, item.ID).Error; spare != nil {
						c.JSON(http.StatusBadRequest, gin.H{
							"Error": "Spare Part not found",
							"Item":  item,
						})
						return
					}

					if sparePart.Quantity < item.Quantity {
						c.JSON(http.StatusBadRequest, gin.H{
							"Error": "Insufficient stock",
							"Item":  item,
						})
						return
					}
					sparePart.Quantity -= item.Quantity
					if quantityUpdate := initializers.DB.Save(&sparePart).Error; quantityUpdate != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"Error": "Failed to update spare part quantity",
							"Item":  item,
						})
						return
					}
				}

				ReportItem := models.SalesReportItems{
					StoreItemsID:      uint(item.ID),
					ItemName:          item.Item,
					Quantity:          item.Quantity,
					Price:             item.Price,
					Category:          item.Category,
					CategoryMachineID: uint(item.CategoryItemsId),
					SalesReportID:     Report.SalesReportID,
				}
				initializers.DB.Create(&ReportItem)
			}
			c.JSON(http.StatusOK, gin.H{
				"Succes": "Succes Create Sales Report",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": create.Error,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Item Not Found",
			"Data":  Sales,
		})
	}
}

func GetSalesReport(c *gin.Context) {
	var salesReport []models.SalesReports

	initializers.DB.Preload("SalesReportItems.CategoryMachine").Find(&salesReport)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Sales Report",
		"Data":   salesReport,
	})
}
