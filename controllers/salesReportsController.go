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

		tx := initializers.DB.Begin()
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Failed to start transaction",
			})
			return
		}

		totalPrice := 0
		for _, item := range Sales.Item {
			totalPrice += item.Price * item.Quantity
		}

		Report := models.SalesReports{
			Date:       dateTime,
			TotalPrice: totalPrice,
		}

		if salesReport := tx.Create(&Report).Error; salesReport != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Failed to create sales report",
			})
			return
		}

		for _, item := range Sales.Item {
			if item.Category == "mesin" {
				var storeItem models.StoreItems
				if store := tx.First(&storeItem, item.ID).Error; store != nil {
					tx.Rollback()
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": "Store item not found",
						"Item":  item,
					})
					return
				}

				if storeItem.Quantity < item.Quantity {
					tx.Rollback()
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": "Insufficient stock",
						"Item":  item,
					})
					return
				}
				storeItem.Quantity -= item.Quantity
				if quantityUpdate := tx.Save(&storeItem).Error; quantityUpdate != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"Error": "Failed to update store item quantity",
						"Item":  item,
					})
					return
				}
			} else if item.Category == "spare_part" {
				var sparePart models.SparePart
				if spare := tx.First(&sparePart, item.ID).Error; spare != nil {
					tx.Rollback()
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": "Spare part not found",
						"Item":  item,
					})
					return
				}

				if sparePart.Quantity < item.Quantity {
					tx.Rollback()
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": "Insufficient stock",
						"Item":  item,
					})
					return
				}
				sparePart.Quantity -= item.Quantity
				if quantityUpdate := tx.Save(&sparePart).Error; quantityUpdate != nil {
					tx.Rollback()
					c.JSON(http.StatusInternalServerError, gin.H{
						"Error": "Failed to update spare part quantity",
						"Item":  item,
					})
					return
				}
			}

			ReportItem := models.SalesReportItems{
				StoreItemsID:  uint(item.ID),
				ItemName:      item.Item,
				Quantity:      item.Quantity,
				Price:         item.Price,
				Category:      item.Category,
				CategoryID:    uint(item.CategoryItemsId),
				SalesReportID: Report.SalesReportID,
			}
			if reportItems := tx.Create(&ReportItem).Error; reportItems != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Failed to create sales report item",
					"Item":  item,
				})
				return
			}
		}

		if commitTransaction := tx.Commit().Error; commitTransaction != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Failed to commit transaction",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Success": "Successfully created Sales Report",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Item Not Found",
			"Data":  Sales,
		})
	}
}

func GetSalesReport(c *gin.Context) {
	var salesReport []models.SalesReports

	initializers.DB.Preload("SalesReportItems.Categories").Order("date DESC").Find(&salesReport)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Sales Report",
		"Data":   salesReport,
	})
}

func DeletedSalesReport(c *gin.Context) {
	id := c.Param("id")
	var salesReport []models.SalesReports
	var salesReportItems []models.SalesReportItems
	sales := initializers.DB.Where("sales_report_id = ?", id).Delete(&salesReportItems)
	if sales.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to Delete Sales Report Items",
		})
		return
	} else {
		initializers.DB.Where("sales_report_id = ?", id).Delete(&salesReport)
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Deleting Sales Report",
		})
	}
}
