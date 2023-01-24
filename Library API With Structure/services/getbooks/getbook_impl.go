package getbooks

import (
	_ "github.com/lib/pq"
	"github.com/Library-RestAPI/models"
	"github.com/Library-RestAPI/repositories/books"
)

type getBookService struct{}

var GetBookService getBookService = getBookService{}

func (getBookService) GetBooks() ([]models.Book, error) {
	// repository
	books, err := books.BookRepositoryImpl.GetBooks()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (getBookService) GetBook(id string) (models.Book, error) {
	// repository
	book, err := books.BookRepositoryImpl.GetBook(id)
	if err != nil {
		return book, err
	}
	return book, nil
}