package checkinsandcheckouts

import (
	"net/http"

	"github.com/Library-RestAPI/initializers"
	"github.com/Library-RestAPI/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CheckInBook(c *gin.Context) (models.Book) {
	//param 
	id := c.Param("id")
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
	// update the quantity
	tempBook.Quantity = tempBook.Quantity + 1
	_, err = initializers.Db.Exec("UPDATE books SET quantity = $1 WHERE id = $2", tempBook.Quantity, tempBook.ID)
	CheckError(err)
	c.IndentedJSON(http.StatusOK, tempBook)	
	return tempBook
}
	
func CheckOutBook(c *gin.Context) (models.Book) {
	//param 
	id := c.Param("id")
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

	// check if the book is available or not
	if tempBook.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not available"})
		return tempBook
	}

	// update the quantity
	tempBook.Quantity = tempBook.Quantity - 1
	_, err = initializers.Db.Exec("UPDATE books SET quantity = $1 WHERE id = $2", tempBook.Quantity, tempBook.ID)
	CheckError(err)
	c.IndentedJSON(http.StatusOK, tempBook)
	return tempBook
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
