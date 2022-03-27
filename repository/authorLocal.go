package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/parvineyvazov/book-api/model"

	"github.com/google/uuid"
)

// GetAuthors get all authors from DB(local)
func (*RepositoryLocal) GetAuthors() (model.Authors, error) {
	fmt.Println("get all authors")

	return Authors, nil
}

// GetAuthor get author by id from DB(local)
func (*RepositoryLocal) GetAuthor(id string) (model.Author, error) {
	fmt.Println("get author with id", id)

	for _, author := range Authors {
		if author.Id == id {
			return author, nil
		}
	}

	return model.Author{}, errors.New("ERROR: Not any author with this id")
}

// SearchAuthor get authors with related search text from DB(local)
func (*RepositoryLocal) SearchAuthors(search_text string) (model.Authors, error) {
	fmt.Println("get author with search text", search_text)

	var authors model.Authors = model.Authors{}

	for _, author := range Authors {
		if strings.Contains(author.Name, search_text) ||
			strings.Contains(author.Surname, search_text) {
			authors = append(authors, author)
		}
	}

	authors = integrateBooks(authors)

	return authors, nil
}

// CreateAuthor create author in DB(local)
func (*RepositoryLocal) CreateAuthor(author model.Author) (model.Author, error) {
	fmt.Println("create author", author)

	author.Id = uuid.New().String()

	Authors = append(Authors, author)

	return author, nil
}

// UpdateAuthor updates author in DB(local)
func (*RepositoryLocal) UpdateAuthor(id string, author model.Author) (model.Author, error) {
	fmt.Println("update author", id, "with", author)

	for index := range Authors {
		if Authors[index].Id == id {
			Authors[index] = author
			Authors[index].Id = id
			return Authors[index], nil
		}
	}

	return model.Author{}, errors.New("ERROR: Not any author with this id")
}

// DeleteAuthor deletes author by id from DB(local)
func (*RepositoryLocal) DeleteAuthor(id string) (model.Author, error) {
	fmt.Println("delete author", id)

	for index, author := range Authors {
		if author.Id == id {
			Authors = append(Authors[:index], Authors[index+1:]...)

			deleteBooksOfAuthor(author.Id)

			return author, nil
		}
	}

	return model.Author{}, errors.New("ERROR: Not any author with this id")
}

// HELPER functions
func integrateBooks(authors model.Authors) model.Authors {

	for author_index, author := range authors {

		for _, book := range Books {

			for _, authorID := range book.AuthorIDs {

				if authorID == author.Id {
					authors[author_index].Books = append(authors[author_index].Books, book)
				}

			}

		}
	}

	return authors
}

func deleteBooksOfAuthor(authorId string) {

	var new_books model.Books = model.Books{}

	for _, book := range Books {

		if !bookContainsAuthor(book, authorId) {
			new_books = append(new_books, book)
		}

	}

	Books = new_books
}

func bookContainsAuthor(book model.Book, authorId string) bool {

	for _, bookAuthorId := range book.AuthorIDs {
		if bookAuthorId == authorId {
			return true
		}
	}

	return false
}
