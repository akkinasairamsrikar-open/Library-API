package checkinsandouts

import (
	"github.com/gin-gonic/gin"
	"github.com/Library-RestAPI/services/checkinsandouts"
)


func CheckInBook(c *gin.Context) {
	checkinsandouts.CheckInBookService.CheckInBook(c)
}


func CheckOutBook(c *gin.Context) {
	checkinsandouts.CheckOutBookService.CheckOutBook(c)
}



	