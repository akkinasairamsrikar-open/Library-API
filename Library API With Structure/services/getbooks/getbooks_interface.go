package getbooks

import (
	"github.com/gin-gonic/gin"
)

type getBookInterface interface {
	GetBooks(c *gin.Context)
	GetBookbyId(c *gin.Context)
}