package main

import (
	"github.com/Library-RestAPI/initializers"
	"github.com/gin-gonic/gin"
	"github.com/Library-RestAPI/controllers")

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDataBase()
}

func main() {
	r := gin.Default()
	r.GET("/books", controllers.GetBooksfromDb)
	r.GET("/books/:id", controllers.GetBookbyIdfromDB)
	r.POST("/books", controllers.CreateBookintoDb)
	r.PATCH("/books/checkin/:id", controllers.CheckInBookfromDb)
	r.PATCH("/books/checkout/:id", controllers.CheckoutBookfromDb)
	r.Run()
}

