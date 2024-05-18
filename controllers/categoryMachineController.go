package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategoryMachine(c *gin.Context) {
	var categoryMachine []models.CategoryMachine

	initializers.DB.Preload("StoreItems").Find(&categoryMachine)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Category Machine",
		"Data":   categoryMachine,
	})
}

func CategoryMachine(c *gin.Context) {
	var categoryMachine struct {
		MachineName string `json:"machine_name"`
	}

	if err := c.BindJSON(&categoryMachine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Machine := models.CategoryMachine{
		CategoryMachineName: categoryMachine.MachineName,
	}

	create := initializers.DB.Create(&Machine).Preload("StoreItems").Find(&Machine)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Category Machine",
			"Data":   Machine,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}
