package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSparePart(c *gin.Context) {
	var sparePart []models.SparePart

	initializers.DB.Preload("Category").Find(&sparePart)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Spare Part",
		"Data":   sparePart,
	})
}

func PreOrderSparePart(c *gin.Context) {
	var sparePart struct {
		SparePartID int `json:"spare_part_id"`
		Quantity    int `json:"quantity"`
		Price       int `json:"price"`
	}

	if err := c.BindJSON(&sparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var spare models.SparePart
	if err := initializers.DB.Where("spare_part_id = ?", sparePart.SparePartID).First(&spare).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Spare Part not found"})
		return
	}

	if spare.Price != sparePart.Price && sparePart.Price != 0 {
		newPrice := sparePart.Price * sparePart.Quantity
		oldPrice := spare.Price * spare.Quantity
		totalQuantity := spare.Quantity + sparePart.Quantity
		spare.Price = (newPrice + oldPrice) / totalQuantity
		spare.Quantity = totalQuantity
	} else {
		spare.Quantity = spare.Quantity + sparePart.Quantity
	}

	if err := initializers.DB.Save(&spare).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Spare Part"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Spare Part Pre Order successfully",
	})
}

func SparePart(c *gin.Context) {
	var sparePart struct {
		SparePartName string `json:"spare_part_name"`
		Quantity      int    `json:"quantity"`
		Price         int    `json:"price"`
		CategoryID    uint   `json:"category_id"`
	}

	if err := c.BindJSON(&sparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	items := models.SparePart{
		SparePartName: sparePart.SparePartName,
		Quantity:      sparePart.Quantity,
		CategoryID:    sparePart.CategoryID,
		Price:         sparePart.Price,
	}

	create := initializers.DB.Create(&items).Preload("Category").Find(&items)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Spare Part",
			"Data":   items,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}

func EditSparePart(c *gin.Context) {
	var sparePart struct {
		SparePartID   uint   `json:"spare_part_id"`
		SparePartName string `json:"spare_part_name"`
		Quantity      int    `json:"quantity"`
		Price         int    `json:"price"`
		CategoryID    uint   `json:"category_id"`
	}

	if err := c.BindJSON(&sparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var spare models.SparePart
	if err := initializers.DB.Where("spare_part_id = ?", sparePart.SparePartID).First(&spare).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Spare Part not found"})
		return
	}

	if sparePart.SparePartName != "" && sparePart.SparePartName != spare.SparePartName {
		spare.SparePartName = sparePart.SparePartName
	}

	if sparePart.Quantity != spare.Quantity {
		spare.Quantity = sparePart.Quantity
	}

	if sparePart.Price != 0 && sparePart.Price != spare.Price {
		spare.Price = sparePart.Price
	}

	if err := initializers.DB.Save(&spare).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Spare Part"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "Spare Part updated successfully",
	})
}

func DeleteSparePart(c *gin.Context) {
	id := c.Param("id")

	var spare models.SparePart

	result := initializers.DB.Where("spare_part_id = ?", id).Find(&spare)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Terjadi kesalahan dalam mencari Spare Part.",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Spare Part yang ingin dihapus tidak ditemukan",
		})
		return
	}

	initializers.DB.Where("spare_part_id = ?", id).Delete(&spare)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Spare Part telah terhapus",
	})
}
