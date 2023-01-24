package create

import (
	"github.com/Library-RestAPI/models"
	"github.com/Library-RestAPI/initializers"
	_ "github.com/lib/pq"
)

func CreateBook(newBook models.Book) (models.Book) {
	// insert the newBook into the database
	_, err := initializers.Db.Exec("INSERT INTO books (id, title, author, quantity) VALUES ($1, $2, $3, $4)", newBook.ID, newBook.Title, newBook.Author, newBook.Quantity)
	CheckError(err)
	return newBook
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
