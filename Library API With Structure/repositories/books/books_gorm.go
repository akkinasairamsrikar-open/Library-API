package books


import (
	"github.com/Library-RestAPI/models"
	"github.com/Library-RestAPI/initializers"
	_ "github.com/lib/pq"
)

type BookRepository struct{}


var BookRepositoryImpl BookRepository = BookRepository{}

func (BookRepository) GetBooks() ([]models.Book, error) {
	// create a temp to store the books
	var books []models.Book
	// query the database
	rows, err := initializers.Db.Query("SELECT * FROM books")
	if err != nil {
		return books, err
	}
	// iterate over the rows and append the books to the temp
	for rows.Next() {
		var book models.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (BookRepository) GetBook(id string) (models.Book, error) {
	// create a temp to store the book
	var book models.Book
	// query the database
	err := initializers.Db.QueryRow("SELECT * FROM books WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.Quantity)
	if err != nil {
		return book, err
	}
	return book, nil
}



