package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SearchMachine(c *gin.Context) {
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

func SearchSparePart(c *gin.Context) {
	categoriesParam := c.Query("categories")
	nameParam := c.Query("name")

	var sparePart []models.SparePart
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

		query = query.Where("category_spare_part_id IN ?", categories)
	}

	if nameParam != "" {
		query = query.Where("spare_part_name LIKE ?", "%"+nameParam+"%")
	}

	err = query.Preload("CategorySparePart").Find(&sparePart).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sparePart)
}
