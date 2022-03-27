package service

import (
	"github.com/parvineyvazov/book-api/model"
	"github.com/parvineyvazov/book-api/repository"
)

type IService interface {
	// Book
	GetBooks() (model.Books, error)
	GetBook(id string) (model.Book, error)
	SearchBooks(search_text string) (model.Books, error)
	CreateBook(book model.Book) (model.Book, error)
	UpdateBook(id string, book model.Book) (model.Book, error)
	DeleteBook(id string) (model.Book, error)

	// Author
	GetAuthors() (model.Authors, error)
	GetAuthor(id string) (model.Author, error)
	SearchAuthors(search_text string) (model.Authors, error)
	CreateAuthor(author model.Author) (model.Author, error)
	UpdateAuthor(id string, author model.Author) (model.Author, error)
	DeleteAuthor(id string) (model.Author, error)
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) IService {
	return &Service{repository}
}
