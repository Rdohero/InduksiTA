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

func DeletedCategoryMachine(c *gin.Context) {
	id := c.Param("id")

	var machine models.CategoryMachine

	result := initializers.DB.Where("category_machine_id = ?", id).Find(&machine)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Terjadi kesalahan dalam mencari Category Machine.",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Category Machine yang ingin dihapus tidak ditemukan",
		})
		return
	}

	deleted := initializers.DB.Where("category_machine_id = ?", id).Delete(&machine)

	if deleted.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Mohon Untuk Menghapus Machine yang terdapat di category machine sebelum menghapus category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Category Machine telah terhapus",
	})
}

func EditCategoryMachine(c *gin.Context) {
	var machines struct {
		CategoryMachineID   uint   `json:"category_machine_id"`
		CategoryMachineName string `json:"category_machine_name"`
	}

	if err := c.BindJSON(&machines); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var machine models.CategoryMachine
	if err := initializers.DB.Where("category_machine_id = ?", machines.CategoryMachineID).First(&machine).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Machine not found"})
		return
	}

	if machines.CategoryMachineName != "" && machines.CategoryMachineName != machine.CategoryMachineName {
		machine.CategoryMachineName = machines.CategoryMachineName
	}

	if err := initializers.DB.Save(&machine).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Category Machine"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Category Machine updated successfully",
	})
}
