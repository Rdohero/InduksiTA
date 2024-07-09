package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategorySparePart(c *gin.Context) {
	var categorySparePart []models.CategorySparePart

	initializers.DB.Preload("SparePart").Find(&categorySparePart)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Category Spare Part",
		"Data":   categorySparePart,
	})
}

func CategorySparePart(c *gin.Context) {
	var categorySparePart struct {
		SparePartName string `json:"spare_part_name"`
	}

	if err := c.BindJSON(&categorySparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Spare := models.CategorySparePart{
		CategorySparePartName: categorySparePart.SparePartName,
	}

	create := initializers.DB.Create(&Spare).Preload("SparePart").Find(&Spare)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Category Spare Part",
			"Data":   Spare,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func DeletedCategorySparePart(c *gin.Context) {
	id := c.Param("id")

	var spare models.CategorySparePart

	result := initializers.DB.Where("category_spare_part_id = ?", id).Find(&spare)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Terjadi kesalahan dalam mencari Category Spare Part.",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Category Spare Part yang ingin dihapus tidak ditemukan",
		})
		return
	}

	deleted := initializers.DB.Where("category_machine_id = ?", id).Delete(&spare)

	if deleted.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Mohon Untuk Menghapus Spare Part yang terdapat di Category Spare Part sebelum menghapus Category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Category Spare Part telah terhapus",
	})
}

func EditCategorySparePart(c *gin.Context) {
	var sparePart struct {
		CategorySparePartID   uint   `json:"category_spare_part_id"`
		CategorySparePartName string `json:"category_spare_part_name"`
	}

	if err := c.BindJSON(&sparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var spare models.CategorySparePart
	if err := initializers.DB.Where("category_spare_part_id = ?", sparePart.CategorySparePartID).First(&spare).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category Spare Part not found"})
		return
	}

	if sparePart.CategorySparePartName != "" && sparePart.CategorySparePartName != spare.CategorySparePartName {
		spare.CategorySparePartName = sparePart.CategorySparePartName
	}

	if err := initializers.DB.Save(&spare).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Category Spare Part"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Category Spare Part updated successfully",
	})
}
