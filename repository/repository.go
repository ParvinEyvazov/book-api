package repository

import (
	"github.com/parvineyvazov/book-api/model"
)

type IRepository interface {

	// Books
	GetBooks() (model.Books, error)
	GetBook(id string) (model.Book, error)
	SearchBooks(search_text string) (model.Books, error)
	CreateBook(book model.Book) (model.Book, error)
	UpdateBook(id string, book model.Book) (model.Book, error)
	DeleteBook(id string) (model.Book, error)

	// Authors
	GetAuthors() (model.Authors, error)
	GetAuthor(id string) (model.Author, error)
	SearchAuthors(search_text string) (model.Authors, error)
	CreateAuthor(author model.Author) (model.Author, error)
	DeleteAuthor(id string) (model.Author, error)
	UpdateAuthor(id string, author model.Author) (model.Author, error)
}

type RepositoryLocal struct {
}

type RepositoryDB struct {
}

// Repository from local
func NewLocalRepository() IRepository {
	return &RepositoryLocal{}
}

// Repository from DB
func NewDBRepository() IRepository {
	return &RepositoryDB{}
}

var Books model.Books = model.Books{}
var Authors model.Authors = model.Authors{}
