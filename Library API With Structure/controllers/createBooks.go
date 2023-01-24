package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/Library-RestAPI/initializers"
	"github.com/Library-RestAPI/models"
)

var err error


func CreateBookintoDb(c *gin.Context) {
	var newBook models.Book
	if err:= c.BindJSON(&newBook); err != nil {
		return
	}
	_, err = initializers.Db.Exec("INSERT INTO books (id, title, author, quantity) VALUES ($1, $2, $3, $4)", newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	CheckError(err)
	c.IndentedJSON(http.StatusCreated, newBook)
}