package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/parvineyvazov/book-api/model"

	"github.com/google/uuid"
)

// GetBooks get all books from DB(local)
func (*RepositoryLocal) GetBooks() (model.Books, error) {
	fmt.Println("get all books")

	return Books, nil
}

// GetBook get book by id from DB(local)
func (*RepositoryLocal) GetBook(id string) (model.Book, error) {
	fmt.Println("get book with id", id)

	for _, v := range Books {
		if v.Id == id {
			return v, nil
		}
	}

	return model.Book{}, errors.New("ERROR: Not any book with this id")
}

// SearchBooks get books with related search text from DB(local)
func (*RepositoryLocal) SearchBooks(search_text string) (model.Books, error) {
	fmt.Println("search books with: ", search_text)

	var books model.Books = model.Books{}

	for _, book := range Books {
		if strings.Contains(book.Description, search_text) ||
			strings.Contains(book.Title, search_text) ||
			bookContainsAuthorID(&book, search_text) {

			books = append(books, book)
		}
	}

	books = integrateAuthors(books)

	return books, nil
}

// CreateBook creates book in DB(local)
func (*RepositoryLocal) CreateBook(book model.Book) (model.Book, error) {
	fmt.Println("create book", book)

	book.Id = uuid.New().String()

	Books = append(Books, book)

	return book, nil
}

// UpdateBook updates book in DB(local)
func (*RepositoryLocal) UpdateBook(id string, book model.Book) (model.Book, error) {
	fmt.Println("update book", id)

	fmt.Println("new book", book)

	for index := range Books {
		if Books[index].Id == id {
			Books[index] = book
			Books[index].Id = id
			return Books[index], nil
		}
	}

	return model.Book{}, errors.New("ERROR: Not any book with this id")
}

// DeleteBook deletes book from DB(local)
func (*RepositoryLocal) DeleteBook(id string) (model.Book, error) {
	fmt.Println("delete book", id)

	for index, book := range Books {
		if book.Id == id {
			Books = append(Books[:index], Books[index+1:]...)
			return book, nil
		}
	}

	return model.Book{}, errors.New("ERROR: Not any book with this id")
}

// HELPER functions
func bookContainsAuthorID(book *model.Book, id string) bool {

	for _, v := range book.AuthorIDs {
		if strings.Contains(v, id) {
			return true
		}
	}

	return false
}

func integrateAuthors(books model.Books) model.Books {

	for index, book := range books {
		for _, authorID := range book.AuthorIDs {
			books[index].Authors = append(books[index].Authors, getAuthorByID(authorID))
		}
	}

	return books
}

func getAuthorByID(id string) model.Author {

	for _, author := range Authors {
		if author.Id == id {
			return author
		}
	}

	return model.Author{}
}
