package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSparePart(c *gin.Context) {
	var sparePart []models.SparePart

	initializers.DB.Preload("CategorySparePart").Find(&sparePart)

	c.JSON(http.StatusOK, gin.H{
		"Succes": "Succes Getting Spare Part",
		"Data":   sparePart,
	})
}

func SparePart(c *gin.Context) {
	var sparePart struct {
		SparePartName       string `json:"spare_part_name"`
		Quantity            int    `json:"quantity"`
		Price               int    `json:"price"`
		CategorySparePartID uint   `json:"category_spare_part_id"`
	}

	if err := c.BindJSON(&sparePart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	Items := models.SparePart{
		SparePartName:       sparePart.SparePartName,
		Quantity:            sparePart.Quantity,
		CategorySparePartID: sparePart.CategorySparePartID,
		Price:               sparePart.Price,
	}

	create := initializers.DB.Create(&Items).Preload("CategorySparePart").Find(&Items)

	if create.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"Succes": "Succes Create Spare Part",
			"Data":   Items,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": create.Error,
		})
	}
}
