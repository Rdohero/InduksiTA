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
		Date        time.Time `json:"date"`
		UserID      uint      `json:"user_id"`
		Name        string    `json:"name"`
		MachineName string    `json:"machine_name"`
		Complaints  string    `json:"complaints"`
	}

	if err := c.BindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Report := models.ServiceReports{
		Date:        Service.Date,
		Name:        Service.Name,
		MachineName: Service.MachineName,
		Complaints:  Service.Complaints,
		StatusID:    1,
		UserID:      Service.UserID,
	}

	create := initializers.DB.Create(&Report).Preload("Status").Preload("User.Role").Preload("ServiceReportsItems").Find(&Report)

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

type ItemsService struct {
	ID              int    `json:"id"`
	Item            string `json:"item"`
	Price           int    `json:"price"`
	Category        string `json:"category"`
	CategoryItemsId int    `json:"category_items_id"`
	Quantity        int    `json:"quantity"`
}

func EditServiceReport(c *gin.Context) {
	var Service struct {
		ID         uint           `json:"id"`
		Complaints string         `json:"complaints"`
		TotalPrice int            `json:"total_price"`
		Item       []ItemsService `json:"item"`
	}

	if err := c.BindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if Service.Item != nil && len(Service.Item) > 0 {
		tx := initializers.DB.Begin()
		if tx.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "Failed to start transaction",
			})
			return
		}

		var serviceReport models.ServiceReports

		if err := tx.Where("service_report_id = ?", Service.ID).First(&serviceReport).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if Service.Complaints != "" && Service.Complaints != serviceReport.Complaints {
			serviceReport.Complaints = Service.Complaints
		}

		if Service.TotalPrice != 0 && Service.TotalPrice != serviceReport.TotalPrice {
			serviceReport.TotalPrice = Service.TotalPrice
		}

		serviceReport.StatusID = 2

		if err := tx.Save(&serviceReport).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update service data"})
			return
		}

		for _, item := range Service.Item {
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

			ReportItem := models.ServiceReportsItems{
				StoreItemsID:    uint(item.ID),
				ItemName:        item.Item,
				Quantity:        item.Quantity,
				Price:           item.Price,
				Category:        item.Category,
				CategoryID:      uint(item.CategoryItemsId),
				ServiceReportID: Service.ID,
			}
			if reportItems := tx.Create(&ReportItem).Error; reportItems != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Failed to create service report item",
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
			"Success": "Successfully Edit Service Report",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Data Item Not Found",
			"Data":  Service,
		})
	}
}

func GetServiceReport(c *gin.Context) {
	var serviceReport []models.ServiceReports

	initializers.DB.Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Service Report",
		"Data":    serviceReport,
	})
}

func GetServiceReportByUserID(c *gin.Context) {
	id := c.Param("id")
	var serviceReport []models.ServiceReports

	initializers.DB.Where("user_id = ?", id).Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Service Report",
		"Data":    serviceReport,
	})
}
