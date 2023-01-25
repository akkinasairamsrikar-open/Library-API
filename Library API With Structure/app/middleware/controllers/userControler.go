package controllers

import (
	"fmt"

	"github.com/Library-RestAPI/app/middleware/helpers"
	"github.com/Library-RestAPI/initializers"
	"github.com/Library-RestAPI/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyPassword(hashedPassword, password string) (bool,string){
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

func SignUp(c *gin.Context) {
	fmt.Println(c)
	var user models.User
	c.BindJSON(&user)
	fmt.Println(user)
	hashedPassword, _ := HashPassword(user.Password)
	user.Password = string(hashedPassword)
	user.Token, _ = helpers.GenerateToken(user.ID)
	user.RefreshToken, _ = helpers.GenerateRefreshToken(user.ID)
	fmt.Println(user)
	_,err := initializers.Db.Exec("INSERT INTO users (id, firstname, lastname, email, password, token, refreshtoken) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.ID, user.Firstname, user.Lastname, user.Email, user.Password, user.Token, user.RefreshToken)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"message": "User not created !"})
	}
	c.JSON(200, gin.H{"user": user})
}

func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	var dbUser models.User
	err := initializers.Db.QueryRow("SELECT * FROM users WHERE email = $1", user.Email).Scan(&dbUser.ID, &dbUser.Firstname, &dbUser.Lastname, &dbUser.Email, &dbUser.Password, &dbUser.Token, &dbUser.RefreshToken)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	// compare password
	if _, err := VerifyPassword(dbUser.Password, user.Password); err != "" {
		c.JSON(404, gin.H{"message": "Password not match"})
		return
	}
	c.JSON(200, gin.H{"user": dbUser})
}


func GetUsers(c *gin.Context) {
	var users []models.User
	rows, err := initializers.Db.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(404, gin.H{"message": "Users not found"})
		return
	}
	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Token, &user.RefreshToken)
		users = append(users, user)
	}
	c.JSON(200, gin.H{"users": users})
}

func GetUserbyID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := initializers.Db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Token, &user.RefreshToken)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		return
	}
	c.JSON(200, gin.H{"user": user})
}
