package routes

import (
	"github.com/Library-RestAPI/app/middleware/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingroutes *gin.Engine) {
	incomingroutes.POST("/signup", controllers.SignUp)
	incomingroutes.POST("/login", controllers.Login)
}
