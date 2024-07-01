package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStoreItems(c *gin.Context) {
	var storeItems []models.StoreItems

	initializers.DB.Preload("CategoryMachine").Find(&storeItems)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Store Items",
		"Data":   storeItems,
	})
}

func StoreItems(c *gin.Context) {
	var storeItems struct {
		StoreItemsName    string `json:"store_items_name"`
		Quantity          int    `json:"quantity"`
		Price             int    `json:"Price"`
		CategoryMachineID uint   `json:"category_machine_id"`
	}

	if err := c.BindJSON(&storeItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	items := models.StoreItems{
		StoreItemsName:    storeItems.StoreItemsName,
		Quantity:          storeItems.Quantity,
		CategoryMachineID: storeItems.CategoryMachineID,
		Price:             storeItems.Price,
	}

	create := initializers.DB.Create(&items).Preload("CategoryMachine").Find(&items)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Store Items",
			"Data":   items,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func EditStoreItems(c *gin.Context) {
	var storeItems struct {
		StoreItemsID      uint   `json:"store_items_id"`
		StoreItemsName    string `json:"store_items_name"`
		Quantity          int    `json:"quantity"`
		Price             int    `json:"Price"`
		CategoryMachineID uint   `json:"category_machine_id"`
	}

	if err := c.BindJSON(&storeItems); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var items models.StoreItems
	if err := initializers.DB.Where("store_items_id = ?", storeItems.StoreItemsID).First(&items).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Items not found"})
		return
	}

	if storeItems.StoreItemsName != "" && storeItems.StoreItemsName != items.StoreItemsName {
		items.StoreItemsName = storeItems.StoreItemsName
	}

	if storeItems.Quantity != 0 && storeItems.Quantity != items.Quantity {
		items.Quantity = storeItems.Quantity
	}

	if storeItems.Price != 0 && storeItems.Price != items.Price {
		items.Price = storeItems.Price
	}

	if err := initializers.DB.Save(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Items Store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Items Store updated successfully",
	})
}

func DeleteStoreItems(c *gin.Context) {
	id := c.Param("id")

	var items models.StoreItems

	result := initializers.DB.Where("store_items_id = ?", id).Find(&items)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Terjadi kesalahan dalam mencari Items.",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Items yang ingin dihapus tidak ditemukan",
		})
		return
	}

	initializers.DB.Where("store_items_id = ?", id).Delete(&items)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Items telah terhapus",
	})
}
