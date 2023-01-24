package app

import (
		"fmt"
		"github.com/gin-gonic/gin"
		"github.com/Library-RestAPI/initializers"
)

var router *gin.Engine


func StartApplication() {
	fmt.Println("Starting Library API")
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	initializers.LoadEnvVariable()
	initializers.ConnectDataBase()
	mapUrls()
	router.Run(":8080")
}

