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
