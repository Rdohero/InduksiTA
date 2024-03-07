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
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     int    `json:"role"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.Preload("Role").First(&user, "email = ?", body.Email)

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if user.Active == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Email has not been verified",
		})
	} else {
		if errPassword != nil || user.UserID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Email atau Username atau Password Salah",
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
