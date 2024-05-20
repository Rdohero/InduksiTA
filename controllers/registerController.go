package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var body struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Address     string `json:"address"`
		NoHandphone string `json:"no_handphone"`
	}

	if c.Bind(&body) != nil {
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

	if checkPassword == nil {
		user := models.User{
			Username:    body.Username,
			Password:    string(hash),
			Address:     body.Address,
			NoHandphone: body.NoHandphone,
		}
		result := initializers.DB.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Email is already in use",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Status": "Succes!",
		})
	}
}
