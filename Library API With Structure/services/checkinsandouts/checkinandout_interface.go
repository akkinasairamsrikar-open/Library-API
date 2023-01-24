package checkinsandouts

import (
	"github.com/gin-gonic/gin"
)

type CheckInOutInterface interface {
	CheckInBook(c *gin.Context)
	CheckOutBook(c *gin.Context)
}