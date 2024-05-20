package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Search(c *gin.Context) {
	categoriesParam := c.Query("categories")
	nameParam := c.Query("name")

	var storeItems []models.StoreItems
	var err error
	query := initializers.DB

	if categoriesParam != "" {
		categoryIDs := strings.Split(categoriesParam, ",")

		var categories []uint
		for _, id := range categoryIDs {
			var categoryID uint
			fmt.Sscanf(id, "%d", &categoryID)
			categories = append(categories, categoryID)
		}

		query = query.Where("category_machine_id IN ?", categories)
	}

	if nameParam != "" {
		query = query.Where("store_items_name LIKE ?", "%"+nameParam+"%")
	}

	err = query.Preload("CategoryMachine").Find(&storeItems).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, storeItems)
}
