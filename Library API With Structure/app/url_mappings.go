package app

import (
	"fmt"
	"github.com/Library-RestAPI/app/middleware/routes"
	"github.com/Library-RestAPI/controllers/checkinsandouts"
	"github.com/Library-RestAPI/controllers/createbooks"
	"github.com/Library-RestAPI/controllers/getbooks"
	"github.com/Library-RestAPI/app/middleware/auth"
)

func mapUrls() {
	fmt.Println("Mapping Urls")
	routes.AuthRouter(router)
	router.Use(auth.Authenticate)
	router.GET("/books", getbooks.GetBooks)
	router.GET("/books/:id", getbooks.GetBookbyID)
	router.POST("/books", createbooks.CreateBook)
	router.PATCH("/books/checkin/:id", checkinsandouts.CheckInBook)
	router.PATCH("/books/checkout/:id", checkinsandouts.CheckOutBook)
}


