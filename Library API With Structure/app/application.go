package app

import (
	"fmt"
	"github.com/Library-RestAPI/initializers"
	// "github.com/Library-RestAPI/app/middleware/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApplication() {
	fmt.Println("Starting Library API")
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//routes.AuthRouter(router)
	initializers.LoadEnvVariable()
	initializers.ConnectDataBase()
	mapUrls()
	router.Run(":8080")
}
