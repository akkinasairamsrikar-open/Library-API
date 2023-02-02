package auth

import (
	"fmt"

	"github.com/Library-RestAPI/app/middleware/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if _, err := helpers.Validatetoken(clientToken); err != "" {
		c.JSON(401, gin.H{"message": err})
		fmt.Println(err)
		c.Abort()
	}
	c.Next()
}
