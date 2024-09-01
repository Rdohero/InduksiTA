package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"path/filepath"
)

func Register(c *gin.Context) {
	var body struct {
		Username    string `form:"username" json:"username"`
		Password    string `form:"password" json:"password"`
		Address     string `form:"address" json:"address"`
		NoHandphone string `form:"no_handphone" json:"no_handphone"`
		Role        uint   `form:"role" json:"role"`
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to read body",
		})

		return
	}

	checkPassword := checkPasswordCriteria(body.Password)
	if checkPassword != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": checkPassword.Error(),
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})

		return
	}

	allowedMIMETypes := []string{"image/jpeg", "image/png", "image/svg"}

	if !IsValidMIMEType(file, allowedMIMETypes) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya menerima jpeg, png, dan svg"})
		return
	}

	if checkPassword == nil {
		var username []models.User

		initializers.DB.Where("username = ? && is_deleted = ?", body.Username, 1).Find(&username)

		if len(username) != 0 {
			initializers.DB.Model(&username).Where("username = ?", body.Username).Update("is_deleted", 0)
		} else {
			basePath := filepath.Join("images", file.Filename)
			os.MkdirAll("images", os.ModePerm)
			filePath := generateUniqueFileName(basePath)
			c.SaveUploadedFile(file, filePath)
			user := models.User{
				Username:    body.Username,
				Password:    string(hash),
				Address:     body.Address,
				NoHandphone: body.NoHandphone,
				Image:       filePath,
				RoleID:      body.Role,
			}
			result := initializers.DB.Create(&user)

			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": result.Error,
				})
				os.Remove(filePath)

				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"Status": "Succes!",
		})
	}
}
