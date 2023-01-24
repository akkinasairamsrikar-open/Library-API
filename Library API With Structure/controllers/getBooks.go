package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/Library-RestAPI/initializers"
	"github.com/Library-RestAPI/models"
)


func GetBooksfromDb(c *gin.Context) {
	
	//get data from database
	//return data
	rows, err := initializers.Db.Query("SELECT * FROM books")
	CheckError(err)
	defer rows.Close()

	// store data in books struct
	var books []models.Book
	for rows.Next() {
		var id string
		var title string
		var author string
		var quantity int
		err = rows.Scan(&id, &title, &author, &quantity)
		CheckError(err)
		books = append(books, models.Book{ID: id, Title: title, Author: author, Quantity: quantity})
	}

	c.IndentedJSON(http.StatusOK, books)
}

func GetBookbyIdfromDB(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")
	// get data from database
	// return data
	rows, err := initializers.Db.Query("SELECT * FROM books WHERE id = $1", id)
	CheckError(err)

	var tempBook models.Book
	for rows.Next() {
		var id string
		var title string
		var author string
		var quantity int
		err = rows.Scan(&id, &title, &author, &quantity)
		CheckError(err)
		tempBook = models.Book{ID: id, Title: title, Author: author, Quantity: quantity}
	}
	c.IndentedJSON(http.StatusOK, tempBook)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
