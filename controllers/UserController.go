package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
)

var saltChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func ValidateCredentials(c *gin.Context) {
	var creds models.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	md5PassSalt := getMD5(creds.Pass, user.Salt)

	if md5PassSalt == user.PassMD5 {
		c.JSON(http.StatusOK, true)
		return
	}

	c.JSONP(http.StatusBadRequest, "Invalid Credentials")
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input models.Credentials
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(&input)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Generate Salt
	salt := randSeq(12)

	//Generate MD5
	md5PassSalt := getMD5(input.Pass, salt)

	// Create user
	user := models.User{
		Model:   gorm.Model{},
		ID:      0,
		Email:   input.Email,
		PassMD5: md5PassSalt,
		Salt:    salt,
		IsAdmin: false,
	}

	models.DB.Create(&user)

	//return the response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func getMD5(pass string, salt string) string {
	passSalt := []byte(pass + salt)
	return fmt.Sprintf("%x", md5.Sum(passSalt))
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = saltChars[rand.Intn(len(saltChars))]
	}
	return string(b)
}
