package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

func ResendOtpEmailPassVer(c *gin.Context) {
	var Resend struct {
		Username string `json:"username"`
	}

	c.Bind(&Resend)

	var Email = "athillahaidar@gmail.com"

	var _, err1 = getUserByUsername(Resend.Username)

	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "User not found",
		})
		return
	}

	otp := rand.Intn(900000) + 100000

	otpStr := fmt.Sprintf("%06d", otp)

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": Resend.Username,
		"otp":      otp,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	}).SignedString([]byte(os.Getenv("SECRET")))

	SimpanOtp(otpStr, token)

	subject := "Email Verificaion"
	HTMLbody :=
		`<html>
			<h1>Code to Verify Email For User : ` + Resend.Username + `</h1>
			<p>` + otpStr + `</p>
		</html>`
	to := []string{Email}
	cc := []string{Email}
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

	var Email2 = "rojaridho8888@gmail.com"

	subject2 := "Email Verificaion"
	HTMLbody2 :=
		`<html>
			<h1>Code to Verify Email For User : ` + Resend.Username + `</h1>
			<p>` + otpStr + `</p>
		</html>`
	to2 := []string{Email2}
	cc2 := []string{Email2}
	// SMTP - Simple Mail Transfer Protocol
	host2 := "smtp.gmail.com"
	port2 := "587"
	address2 := host2 + ":" + port2
	// Set up authentication information
	auth2 := smtp.PlainAuth("", "itemsgates@gmail.com", "supg zioz tclu ewdx", host2)
	// Construct the message
	msg2 := []byte(
		"From: " + "Items Gate <itemsgates@gmail.com>" + "\n" +
			"To: " + strings.Join(to2, ",") + "\n" +
			"Cc: " + strings.Join(cc2, ",") + "\n" +
			"Subject: " + subject2 + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-s8\"\r\n" +
			"\r\n" +
			HTMLbody2)
	err2 := smtp.SendMail(address2, auth2, "Items Gate", to2, msg2)

	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Error",
			"Error":  err2.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Send Code Succes",
	})
}

func ForgotPassword(c *gin.Context) {
	var Otp struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Otp      string `json:"otp"`
	}
	fmt.Println(Otp.Username)
	fmt.Println("belum di bind")
	c.BindJSON(&Otp)

	var token2, err1 = DapatkanOtpString(Otp.Otp)

	if token2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err1.Error(),
		})
		return
	}

	token, _ := jwt.Parse(token2, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			HapusOtp(Otp.Otp)
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Otp Has Been Expired",
			})
			return
		}

		user, err := getUserByUsername(claims["username"].(string))

		var userUpdate []models.User

		if Otp.Username == user.Username {
			if Otp.Otp == strconv.Itoa(int(claims["otp"].(float64))) {

				HapusOtp(Otp.Otp)

				hash, hashErr := bcrypt.GenerateFromPassword([]byte(Otp.Password), 14)

				if hashErr != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"Error": "Failed to hash password",
					})

					return
				}

				initializers.DB.First(&userUpdate, user.UserID).Update("password", string(hash))

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{
						"Error": "Please try verification email again",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"Status": "Succes",
				})
				return
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"Error": "Otp Not Valid",
				})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Otp Not Valid",
			})
		}
	} else {
		HapusOtp(Otp.Otp)
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Otp Has Been Expired",
		})
		return
	}
}

func getUserByUsername(username string) (*models.User, error) {
	var u models.User
	result := initializers.DB.Where("username = ?", username).First(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return &u, nil
}
