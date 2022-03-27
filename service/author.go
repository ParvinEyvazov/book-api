package service

import (
	"github.com/parvineyvazov/book-api/model"
)

func (s *Service) GetAuthors() (model.Authors, error) {

	authors, err := s.repository.GetAuthors()

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *Service) GetAuthor(id string) (model.Author, error) {

	author, err := s.repository.GetAuthor(id)

	if err != nil {
		return model.Author{}, err
	}

	return author, err
}

func (s *Service) SearchAuthors(search_text string) (model.Authors, error) {

	authors, err := s.repository.SearchAuthors(search_text)

	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *Service) CreateAuthor(author model.Author) (model.Author, error) {

	created_author, err := s.repository.CreateAuthor(author)

	if err != nil {
		return model.Author{}, err
	}

	return created_author, nil
}

func (s *Service) UpdateAuthor(id string, author model.Author) (model.Author, error) {

	updated_author, err := s.repository.UpdateAuthor(id, author)

	if err != nil {
		return model.Author{}, err
	}

	return updated_author, nil
}

func (s *Service) DeleteAuthor(id string) (model.Author, error) {
	author, err := s.repository.DeleteAuthor(id)

	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}
