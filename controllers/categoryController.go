package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategory(c *gin.Context) {
	var category []models.Category

	initializers.DB.Preload("StoreItems").Preload("SparePart").Find(&category)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Category Machine",
		"Data":   category,
	})
}

func CategoryPost(c *gin.Context) {
	var category struct {
		CategoryName string `json:"category_name"`
	}

	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Machine := models.Category{
		CategoryName: category.CategoryName,
	}

	create := initializers.DB.Create(&Machine).Preload("StoreItems").Preload("SparePart").Find(&Machine)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Category",
			"Data":   Machine,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func DeletedCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category

	result := initializers.DB.Where("category_id = ?", id).Find(&category)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Terjadi kesalahan dalam mencari Category.",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Category yang ingin dihapus tidak ditemukan",
		})
		return
	}

	deleted := initializers.DB.Where("category_id = ?", id).Delete(&category)

	if deleted.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Mohon Untuk Menghapus Machine / Spare Part yang terdapat di category sebelum menghapus category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Category telah terhapus",
	})
}

func EditCategory(c *gin.Context) {
	var categories struct {
		CategoryID   uint   `json:"category_id"`
		CategoryName string `json:"category_name"`
	}

	if err := c.BindJSON(&categories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var category models.Category
	if err := initializers.DB.Where("category_id = ?", categories.CategoryName).First(&category).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if categories.CategoryName != "" && categories.CategoryName != category.CategoryName {
		category.CategoryName = categories.CategoryName
	}

	if err := initializers.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Category updated successfully",
	})
}
