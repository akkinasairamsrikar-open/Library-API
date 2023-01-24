package getbooks

import (
	"github.com/gin-gonic/gin"
	"github.com/Library-RestAPI/services/getbooks"
)


func GetBooks(c *gin.Context) {
	// service
	books, err := getbooks.GetBookService.GetBooks()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, books)
}

func GetBookbyID(c *gin.Context) {
	// service
	book, err := getbooks.GetBookService.GetBook(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, book)
}



