package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

func ResendOtpEmailPassVer(c *gin.Context) {
	var Resend struct {
		Email string
	}

	c.Bind(&Resend)

	var _, err1 = getUserByEmail(Resend.Email)

	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": "User not found",
		})
		return
	}

	otp := rand.Intn(900000) + 100000

	otpStr := fmt.Sprintf("%06d", otp)

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": Resend.Email,
		"otp":   otp,
		"exp":   time.Now().Add(time.Minute * 1).Unix(),
	}).SignedString([]byte(os.Getenv("SECRET")))

	SimpanOtp(otpStr, token)

	subject := "Email Verificaion"
	HTMLbody :=
		`<html>
			<h1>Code to Verify Email</h1>
			<p>` + otpStr + `</p>
		</html>`
	to := []string{Resend.Email}
	cc := []string{Resend.Email}
	// SMTP - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// Set up authentication information
	auth := smtp.PlainAuth("", "itemsgates@gmail.com", "supg zioz tclu ewdx", host)
	// Construct the message
	msg := []byte(
		"From: " + "Items Gate <itemsgates@gmail.com>" + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Cc: " + strings.Join(cc, ",") + "\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-s8\"\r\n" +
			"\r\n" +
			HTMLbody)
	err := smtp.SendMail(address, auth, "Items Gate", to, msg)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Error",
			"Error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Resend Code Succes",
	})
}

func getUserByEmail(email string) (*models.User, error) {
	var u models.User
	result := initializers.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}
