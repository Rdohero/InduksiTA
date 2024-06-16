package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
)

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Failed Reading Body",
			"Error":   err.Error(),
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if errPassword != nil || user.UserID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Username atau Password Salah",
		})
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.UserID,
		})

		tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
		c.JSON(http.StatusOK, gin.H{
			"Token": tokenString,
		})
	}
}

func LoginOwner(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Failed Reading Body",
			"Error":   err.Error(),
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if user.RoleID != 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bukan Akun Owner",
		})
	} else {
		if errPassword != nil || user.UserID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Username atau Password Salah",
			})
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user.UserID,
			})

			tokenString, _ := token.SignedString([]byte(os.Getenv("SECRET")))
			c.JSON(http.StatusOK, gin.H{
				"Token": tokenString,
			})
		}
	}
}
