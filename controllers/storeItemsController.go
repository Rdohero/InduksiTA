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

	Items := models.StoreItems{
		StoreItemsName:    storeItems.StoreItemsName,
		Quantity:          storeItems.Quantity,
		CategoryMachineID: storeItems.CategoryMachineID,
		Price:             storeItems.Price,
	}

	create := initializers.DB.Create(&Items).Preload("CategoryMachine").Find(&Items)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Store Items",
			"Data":   Items,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}
