package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/parvineyvazov/book-api/model"
)

// GetBooks get all books from DB
func (*RepositoryDB) GetBooks() (model.Books, error) {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM Book")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var books model.Books
	for results.Next() {
		fmt.Println("next")
		var book model.Book
		// for each row, scan the result into our tag composite object

		var authorInfo []uint8

		err = results.Scan(&book.Id, &book.Title, &book.Description, &book.Publication_date, &authorInfo)
		if err != nil {
			panic(err.Error())
		}

		book.AuthorIDs = getAuthorArray(authorInfo)

		books = append(books, book)
	}

	fmt.Println(`books`, books)

	for _, v := range books {
		fmt.Println("author ", v)
	}

	return books, nil
}

// GetBook get book by id from DB
func (*RepositoryDB) GetBook(id string) (model.Book, error) {

	return model.Book{}, nil
}

// SearchBooks get books with related search text from DB
func (*RepositoryDB) SearchBooks(search_text string) (model.Books, error) {

	return nil, nil
}

// CreateBook creates book in DB
func (*RepositoryDB) CreateBook(book model.Book) (model.Book, error) {

	return model.Book{}, nil
}

// UpdateBook updates book in DB
func (*RepositoryDB) UpdateBook(id string, book model.Book) (model.Book, error) {

	return model.Book{}, nil
}

// DeleteBook deletes book from DB
func (*RepositoryDB) DeleteBook(id string) (model.Book, error) {

	return model.Book{}, nil
}

// HELPER functions
func getAuthorArray(authorBytes []uint8) []string {

	return []string{}
}
