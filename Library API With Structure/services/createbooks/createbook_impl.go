package createbooks

import (
	_ "github.com/lib/pq"
	"github.com/Library-RestAPI/models"
	"github.com/Library-RestAPI/repositories/create"
)

type createBookService struct{}


var CreateBookService createBookService = createBookService{}

func (createBookService) CreateBook(newBook models.Book) (models.Book) {
	// insert the newBook into the database
	// call repository
	newBook = create.CreateBook(newBook)
	return newBook
}


