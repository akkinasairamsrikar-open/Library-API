package checkinsandouts

import (
	"github.com/Library-RestAPI/repositories/checkinsandcheckouts"
	"github.com/gin-gonic/gin"
)

// CheckInBook checks in a book

type checkInBookService struct{}

var CheckInBookService checkInBookService = checkInBookService{}

// take c as a parameter
func (checkInBookService) CheckInBook(c *gin.Context) {
	// call the repository function
	checkinsandcheckouts.CheckInBook(c)
}

// CheckOutBook checks out a book
type checkOutBookService struct{}

var CheckOutBookService checkOutBookService = checkOutBookService{}

// take c as a parameter
func (checkOutBookService) CheckOutBook(c *gin.Context) {
	// call the repository function
	checkinsandcheckouts.CheckOutBook(c)
}
