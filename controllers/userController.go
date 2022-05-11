package controllers

import (
	"crypto/sha512"
	"encoding/hex"
	"github.com/dmnyu/go-medialog/database"
	"github.com/dmnyu/go-medialog/models"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func GetUsers(c *gin.Context) {
	users, err := database.FindUsers()
	if err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "users-show.html", gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {}

func NewUser(c *gin.Context) {
	c.HTML(http.StatusOK, "users-new.html", gin.H{})
}

func CreateUser(c *gin.Context) {
	var createUser = models.CreateUser{}
	if err := c.Bind(&createUser); err != nil {
		log.Printf("[ERROR] [APP] %s", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if createUser.Password1 != createUser.Password2 {
		log.Println("[ERROR] [APP] passwords do not match")
		c.JSON(http.StatusBadRequest, "passwords do not match")
		return
	}

	user := models.User{}
	user.FirstName = createUser.FirstName
	user.LastName = createUser.LastName
	user.Email = createUser.Email
	user.IsAdmin = createUser.IsAdmin
	user.Salt = randomStringRunes(16)
	user.PassMD5 = getSHA256Hash(createUser.Password1 + user.Salt)

	if err := database.CreateUser(&user); err != nil {
		log.Printf("[ERROR] [DATABASE] %s", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func EditUser(c *gin.Context) {}

func UpdateUser(c *gin.Context) {}

func DeleteUser(c *gin.Context) {}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()_+{}[]:;<>,.?/")

func randomStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getSHA256Hash(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}
