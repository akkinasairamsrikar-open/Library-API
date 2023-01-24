package createbooks

import (
	"fmt"
	"net/http"
	"github.com/Library-RestAPI/services/createbooks"
	"github.com/Library-RestAPI/validators"
	"github.com/gin-gonic/gin"
	"github.com/Library-RestAPI/models"
)

var BOOK models.Book

func CreateBook(c *gin.Context) {
	// Validator
	book, err := validators.BookValidatorHandler(c) 
	fmt.Println(book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Service
	BOOK = createbooks.CreateBookService.CreateBook(book)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": BOOK})

	// Step 3: Transformer
	// model (response)
}
