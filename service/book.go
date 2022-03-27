package service

import (
	"github.com/parvineyvazov/book-api/model"
)

func (s *Service) GetBooks() (model.Books, error) {

	books, err := s.repository.GetBooks()

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) GetBook(id string) (model.Book, error) {

	book, err := s.repository.GetBook(id)

	if err != nil {
		return model.Book{}, err
	}

	return book, err
}

func (s *Service) SearchBooks(search_text string) (model.Books, error) {

	books, err := s.repository.SearchBooks(search_text)

	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Service) CreateBook(book model.Book) (model.Book, error) {

	created_book, err := s.repository.CreateBook(book)

	if err != nil {
		return model.Book{}, err
	}

	return created_book, nil
}

func (s *Service) UpdateBook(id string, book model.Book) (model.Book, error) {

	updated_book, err := s.repository.UpdateBook(id, book)

	if err != nil {
		return model.Book{}, err
	}

	return updated_book, nil
}

func (s *Service) DeleteBook(id string) (model.Book, error) {
	book, err := s.repository.DeleteBook(id)

	if err != nil {
		return model.Book{}, err
	}

	return book, nil
}
