package controllers

import (
	"InduksiTA/initializers"
	"InduksiTA/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func GetAllUser(c *gin.Context) {
	var user []models.User
	if err := initializers.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func ChangeProfileUser(c *gin.Context) {
	var requestData struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Address     string `json:"address"`
		NoHandphone string `json:"no_handphone"`
	}

	if err := c.Bind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("user_id = ?", requestData.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if requestData.Password != "" {
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(requestData.Password), 14)
		if hashErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			})
			return
		}
		user.Password = string(hash)
	}

	if requestData.Username != "" && requestData.Username != user.Username {
		user.Username = requestData.Username
	}

	if requestData.Address != "" && requestData.Address != user.Address {
		user.Address = requestData.Address
	}

	if requestData.NoHandphone != "" && requestData.NoHandphone != user.NoHandphone {
		user.NoHandphone = requestData.NoHandphone
	}

	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": "User data updated successfully",
	})
}

func UpdatePhotoProfile(c *gin.Context) {
	id := c.Param("id")

	var user1 []models.User
	initializers.DB.Where("user_id = ?", id).Find(&user1)

	oldfoto := user1[0].Image
	os.Remove(oldfoto)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allowedMIMETypes := []string{"image/jpeg", "image/png", "image/svg"}

	if !IsValidMIMEType(file, allowedMIMETypes) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hanya menerima jpeg, png, dan svg"})
		return
	}
	// Define the path where the file will be saved
	basePath := filepath.Join("images", file.Filename)
	// Create the "images" directory if it doesn't exist
	os.MkdirAll("images", os.ModePerm)
	// Save the file to the defined path
	filePath := generateUniqueFileName(basePath)
	c.SaveUploadedFile(file, filePath)

	user1[0].Image = filePath

	initializers.DB.Model(&user1).Where("user_id = ?", id).Update("image", user1[0].Image)

	c.JSON(http.StatusOK, user1)
}

func IsValidMIMEType(file *multipart.FileHeader, allowedMIMETypes []string) bool {
	src, err := file.Open()
	if err != nil {
		return false
	}
	defer src.Close()

	// Membaca tipe MIME dari file
	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return false
	}

	// Menggunakan http.DetectContentType untuk mendeteksi tipe MIME
	fileType := http.DetectContentType(buffer)

	// Memeriksa apakah tipe MIME ada dalam daftar yang diizinkan
	for _, allowedType := range allowedMIMETypes {
		if fileType == allowedType {
			return true
		}
	}

	return false
}

func generateUniqueFileName(basePath string) string {
	extension := filepath.Ext(basePath)
	name := strings.TrimSuffix(basePath, extension)

	counter := 1
	for {
		newPath := basePath
		if counter > 1 {
			newPath = fmt.Sprintf("%s_%d%s", name, counter, extension)
		}

		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}

		counter++
	}
}

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
