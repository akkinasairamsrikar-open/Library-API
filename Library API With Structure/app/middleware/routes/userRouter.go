package routes

import (
	"github.com/Library-RestAPI/app/middleware/controllers"
	"github.com/Library-RestAPI/app/middleware/auth"
	"github.com/gin-gonic/gin"
)

func UserRouter(incomingroutes *gin.Engine) {
	incomingroutes.Use(auth.Authenticate)
	incomingroutes.GET("/user", controllers.GetUsers)
	incomingroutes.POST("/user", controllers.GetUserbyID)
}