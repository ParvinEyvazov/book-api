package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/parvineyvazov/book-api/mock"
	"github.com/parvineyvazov/book-api/model"
	"github.com/parvineyvazov/book-api/service"
	"github.com/stretchr/testify/assert"
)

func Test_GetBooks(t *testing.T) {

	t.Run("SERVICE [success case] get books", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockBooks := model.Books{
			{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}},
			{Id: "2", Title: "title 2", Description: "desc 2", Publication_date: "2022", AuthorIDs: []string{"2"}},
		}

		mockRepository.
			EXPECT().
			GetBooks().
			Return(mockBooks, nil).
			Times(1)

		service := service.NewService(mockRepository)
		books, err := service.GetBooks()

		assert.Equal(t, mockBooks[0].Id, books[0].Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] get books", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockRepository.
			EXPECT().
			GetBooks().
			Return(nil, errors.New("ERROR: error while fetching books")).
			Times(1)

		service := service.NewService(mockRepository)
		books, err := service.GetBooks()

		assert.Nil(t, books)
		assert.NotNil(t, err)
	})

}

func Test_GetBook(t *testing.T) {

	t.Run("SERVICE [success case] get book", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockBook := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			GetBook(mockID).
			Return(mockBook, nil).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.GetBook(mockID)

		assert.Equal(t, mockBook.Id, book.Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] get books", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"

		mockRepository.
			EXPECT().
			GetBook(mockID).
			Return(model.Book{}, errors.New("ERROR: Not any book with this id")).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.GetBook(mockID)

		assert.Equal(t, model.Book{}, book)
		assert.NotNil(t, err)
	})

}

func Test_SearchBooks(t *testing.T) {

	t.Run("SERVICE [success case] search books", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockSearchText := "test"
		mockBooks := model.Books{
			{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}},
			{Id: "2", Title: "title 2", Description: "desc 2", Publication_date: "2022", AuthorIDs: []string{"2"}},
		}

		mockRepository.
			EXPECT().
			SearchBooks(mockSearchText).
			Return(mockBooks, nil).
			Times(1)

		service := service.NewService(mockRepository)
		books, err := service.SearchBooks(mockSearchText)

		assert.Equal(t, mockBooks[0].Id, books[0].Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] search books", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockSearchText := "test"

		mockRepository.
			EXPECT().
			SearchBooks(mockSearchText).
			Return(nil, errors.New("ERROR: error while searching books")).
			Times(1)

		service := service.NewService(mockRepository)
		books, err := service.SearchBooks(mockSearchText)

		assert.Nil(t, books)
		assert.NotNil(t, err)
	})

}

func Test_CreateBook(t *testing.T) {

	t.Run("SERVICE [success case] create book", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockBook := model.Book{Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			CreateBook(mockBook).
			Return(mockBook, nil).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.CreateBook(mockBook)

		assert.NotNil(t, book)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] create book", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockBook := model.Book{Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			CreateBook(mockBook).
			Return(model.Book{}, errors.New("ERROR: error while creating book")).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.CreateBook(mockBook)

		assert.Equal(t, model.Book{}, book)
		assert.NotNil(t, err)
	})
}

func Test_UpdateBook(t *testing.T) {

	t.Run("SERVICE [success case] update book", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockBook := model.Book{Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			UpdateBook(mockID, mockBook).
			Return(mockBook, nil).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.UpdateBook(mockID, mockBook)

		assert.NotNil(t, book)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] update book", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockBook := model.Book{Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			UpdateBook(mockID, mockBook).
			Return(model.Book{}, errors.New("ERROR: error while updating book")).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.UpdateBook(mockID, mockBook)

		assert.Equal(t, model.Book{}, book)
		assert.NotNil(t, err)
	})

}

func Test_DeleteBook(t *testing.T) {

	t.Run("SERVICE [success case] delete book", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockBook := model.Book{Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockRepository.
			EXPECT().
			DeleteBook(mockID).
			Return(mockBook, nil).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.DeleteBook(mockID)

		assert.NotNil(t, book)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] delete book", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"

		mockRepository.
			EXPECT().
			DeleteBook(mockID).
			Return(model.Book{}, errors.New("ERROR: error while deleting book")).
			Times(1)

		service := service.NewService(mockRepository)
		book, err := service.DeleteBook(mockID)

		assert.Equal(t, model.Book{}, book)
		assert.NotNil(t, err)
	})

}
