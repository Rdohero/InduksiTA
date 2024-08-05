package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ServiceReport(c *gin.Context) {
	var Service struct {
		Date        string `json:"date"`
		UserID      uint   `json:"user_id"`
		Name        string `json:"name"`
		MachineName string `json:"machine_name"`
		Complaints  string `json:"complaints"`
	}

	if err := c.BindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	parsedDate, err := time.Parse("2006-01-02", Service.Date)
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

	Report := models.ServiceReports{
		Date:        dateTime,
		Name:        Service.Name,
		MachineName: Service.MachineName,
		Complaints:  Service.Complaints,
		StatusID:    1,
		UserID:      Service.UserID,
	}

	create := initializers.DB.Create(&Report)

	if create.Error == nil {
		var Reports []models.ServiceReports
		initializers.DB.Preload("Status").Preload("User.Role").Preload("ServiceReportsItems").
			Order("date DESC").
			Order("service_report_id DESC").Find(&Reports)

		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Service Report",
			"Data":   Reports,
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
	}
}

func GetServiceReport(c *gin.Context) {
	var serviceReport []models.ServiceReports

	initializers.DB.Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").
		Order("date DESC").
		Order("service_report_id DESC").
		Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Service Report",
		"Data":    serviceReport,
	})
}

func GetServiceReportByStatusID(c *gin.Context) {
	id := c.Param("id")
	var serviceReport []models.ServiceReports

	initializers.DB.Where("status_id = ?", id).Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").
		Order("date DESC").
		Order("service_report_id DESC").
		Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Service Report",
		"Data":    serviceReport,
	})
}

func GetServiceReportByUserID(c *gin.Context) {
	id := c.Param("id")
	var serviceReport []models.ServiceReports

	initializers.DB.Where("user_id = ? && status_id = ?", id, 1).Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").
		Order("date ASC").
		Order("service_report_id ASC").
		Find(&serviceReport)

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success Getting Service Report",
		"Data":    serviceReport,
	})
}

func GetServiceReportsLastDays(c *gin.Context) {
	daysStr := c.Query("days")
	monthsStr := c.Query("months")
	yearsStr := c.Query("years")

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		days = 0
	}
	months, err := strconv.Atoi(monthsStr)
	if err != nil {
		months = 0
	}
	years, err := strconv.Atoi(yearsStr)
	if err != nil {
		years = 0
	}

	var serviceReports []models.ServiceReports
	adjustedDate := time.Now().AddDate(-years, -months, -days)
	if err := initializers.DB.Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").Where("date >= ?", adjustedDate).Order("date DESC").
		Order("service_report_id DESC").
		Find(&serviceReports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serviceReports)
}

func GetServiceReportsByDateRange(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format. Use YYYY-MM-DD."})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format. Use YYYY-MM-DD."})
		return
	}

	var serviceReports []models.ServiceReports
	if err := initializers.DB.Preload("Status").Preload("User.Role").Preload("ServiceReportsItems.Categories").Where("date BETWEEN ? AND ?", startDate, endDate).Order("date DESC").
		Order("service_report_id DESC").
		Find(&serviceReports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, serviceReports)
}
