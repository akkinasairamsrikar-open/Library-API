package app

import (
	"fmt"
	"github.com/Library-RestAPI/controllers/checkinsandouts"
	"github.com/Library-RestAPI/controllers/createbooks"
	"github.com/Library-RestAPI/controllers/getbooks"
)

func mapUrls(){
	fmt.Println("Mapping Urls")
	
	router.GET("/books", getbooks.GetBooks)
	router.GET("/books/:id", getbooks.GetBookbyID)
	router.POST("/books", createbooks.CreateBook)
	router.PATCH("/books/checkin/:id", checkinsandouts.CheckInBook)
	router.PATCH("/books/checkout/:id", checkinsandouts.CheckOutBook)
}


