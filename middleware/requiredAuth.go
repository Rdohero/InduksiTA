package middleware

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func RequiredAuth(c *gin.Context) {
	var tokenBody struct {
		Token string
	}

	if c.Bind(&tokenBody) != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	tokenString := tokenBody.Token

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.UserID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Expired Login",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
