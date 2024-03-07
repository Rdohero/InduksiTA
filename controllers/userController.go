package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func checkPasswordCriteria(password string) error {
	var err error
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
			err = errors.New("Pa")
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	// check password length
	if 7 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	// create error for any criteria not passed
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdLength || !pswdNoSpaces {
		switch false {
		case pswdLowercase:
			err = errors.New("Password must contain atleast one lower case letter")
		case pswdUppercase:
			err = errors.New("Password must contain atleast one uppercase letter")
		case pswdNumber:
			err = errors.New("Password must contain atleast one number")
		case pswdLength:
			err = errors.New("Passward length must atleast 12 characters and less than 60")
		case pswdNoSpaces:
			err = errors.New("Password cannot have any spaces")
		}
		return err
	}
	return nil
}

func checkEmailValid(email string) error {
	// check email syntax is valid
	//func MustCompile(str string) *Regexp
	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		fmt.Println(err)
		return errors.New("sorry, something went wrong")
	}
	rg := emailRegex.MatchString(email)
	if rg != true {
		return errors.New("Email address is not a valid syntax, please check again")
	}
	// check email length
	if len(email) < 4 {
		return errors.New("Email length is too short")
	}
	if len(email) > 253 {
		return errors.New("Email length is too long")
	}
	return nil
}

func checkEmailDomain(email string) error {
	i := strings.Index(email, "@")
	host := email[i+1:]
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		err = errors.New("Could not find email's domain server, please check and try again")
		return err
	}
	return nil
}

func GetUserById(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User Tidak ditemukan",
		})
	}
	c.JSON(http.StatusOK, user)
}

func ForgotPassword(c *gin.Context) {
	var ForgotPwd struct {
		Password string
		Email    string
		Otp      string
	}
	c.Bind(&ForgotPwd)

	var user models.User

	initializers.DB.Where("email = ?", ForgotPwd.Email).First(&user)

	if user.UserID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Error",
			"Error":  "User not found",
		})
		return
	}

	var tokenString, _ = DapatkanOtpString(ForgotPwd.Otp)

	if tokenString == "" {
		c.JSON(http.StatusOK, gin.H{
			"Status": "Error",
			"Error":  "Otp Not Valid",
		})
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			HapusOtp(ForgotPwd.Otp)
			c.JSON(http.StatusOK, gin.H{
				"Status": "Error",
				"Error":  "Otp Has Been Expired",
			})
			return
		}

		if ForgotPwd.Email == user.Email {
			if ForgotPwd.Otp == strconv.Itoa(int(claims["otp"].(float64))) {
				hash, _ := bcrypt.GenerateFromPassword([]byte(ForgotPwd.Password), 14)

				initializers.DB.First(&user).Update("password", string(hash))

				HapusOtp(ForgotPwd.Otp)

				c.JSON(http.StatusOK, gin.H{
					"Status": "Succes",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"Status": "Error",
					"Error":  "Otp Not Valid",
				})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Status": "Error",
				"Error":  "Otp Not Valid",
			})
		}
	} else {
		HapusOtp(ForgotPwd.Otp)
		c.JSON(http.StatusOK, gin.H{
			"Status": "Error",
			"Error":  "Otp Has Been Expired",
		})
		return
	}
}
