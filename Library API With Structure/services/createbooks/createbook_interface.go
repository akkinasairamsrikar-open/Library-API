package createbooks

import (
	"github.com/gin-gonic/gin"
)

type CreateBookInterface interface {
	CreateBook(c *gin.Context)
}

