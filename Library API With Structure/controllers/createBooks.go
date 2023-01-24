package controllers

import (
	"fmt"
	"net/http"
	"github.com/Library-RestAPI/initializers"
	"github.com/Library-RestAPI/validators"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var err error


func CreateBookintoDb(c *gin.Context) {
	// validate the input from validators/bookValidator.go
	newBook,validation,message := validators.BookValidatorHandler(c)
	if !validation {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}
	fmt.Println(message)
	// insert the newBook into the database
	_, err = initializers.Db.Exec("INSERT INTO books (id, title, author, quantity) VALUES ($1, $2, $3, $4)", newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	CheckError(err)
	c.IndentedJSON(http.StatusCreated, newBook)
}




