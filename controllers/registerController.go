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
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
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

	checkEmail := checkEmailValid(body.Email)
	if checkEmail != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": checkEmail.Error(),
		})
	}

	checkEmailD := checkEmailDomain(body.Email)
	if checkEmailD != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": checkEmailD.Error(),
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})

		return
	}

	if checkPassword == nil && checkEmail == nil && checkEmailD == nil {
		user := models.User{
			Email:    body.Email,
			Username: body.Username,
			Password: string(hash),
			Active:   false,
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
