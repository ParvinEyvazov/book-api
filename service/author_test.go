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

func Test_GetAuthors(t *testing.T) {

	t.Run("SERVICE [success case] get authors", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockAuthors := model.Authors{
			{Id: "1", Name: "author name 1", Surname: "Author surname 1"},
			{Id: "2", Name: "author name 2", Surname: "Author surname 2"},
		}

		mockRepository.
			EXPECT().
			GetAuthors().
			Return(mockAuthors, nil).
			Times(1)

		service := service.NewService(mockRepository)
		authors, err := service.GetAuthors()

		assert.Equal(t, mockAuthors[0].Id, authors[0].Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] get authors", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockRepository.
			EXPECT().
			GetAuthors().
			Return(nil, errors.New("ERROR: error while fetching authors")).
			Times(1)

		service := service.NewService(mockRepository)
		authors, err := service.GetAuthors()

		assert.Nil(t, authors)
		assert.NotNil(t, err)
	})

}

func Test_GetAuthor(t *testing.T) {

	t.Run("SERVICE [success case] get author", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockAuthor := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			GetAuthor(mockID).
			Return(mockAuthor, nil).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.GetAuthor(mockID)

		assert.Equal(t, mockAuthor.Id, author.Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] get author", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"

		mockRepository.
			EXPECT().
			GetAuthor(mockID).
			Return(model.Author{}, errors.New("ERROR: Not any author with this id")).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.GetAuthor(mockID)

		assert.Equal(t, model.Author{}, author)
		assert.NotNil(t, err)
	})

}

func Test_SearchAuthors(t *testing.T) {

	t.Run("SERVICE [success case] search authors", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockSearchText := "test"
		mockAuthors := model.Authors{
			{Id: "1", Name: "author name 1", Surname: "Author surname 1"},
			{Id: "2", Name: "author name 2", Surname: "Author surname 2"},
		}

		mockRepository.
			EXPECT().
			SearchAuthors(mockSearchText).
			Return(mockAuthors, nil).
			Times(1)

		service := service.NewService(mockRepository)
		authors, err := service.SearchAuthors(mockSearchText)

		assert.Equal(t, mockAuthors[0].Id, authors[0].Id)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] search authors", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockSearchText := "test"

		mockRepository.
			EXPECT().
			SearchAuthors(mockSearchText).
			Return(nil, errors.New("ERROR: error while searching authors")).
			Times(1)

		service := service.NewService(mockRepository)
		authors, err := service.SearchAuthors(mockSearchText)

		assert.Nil(t, authors)
		assert.NotNil(t, err)
	})

}

func Test_CreateAuthor(t *testing.T) {

	t.Run("SERVICE [success case] create author", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockAuthor := model.Author{Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			CreateAuthor(mockAuthor).
			Return(mockAuthor, nil).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.CreateAuthor(mockAuthor)

		assert.NotNil(t, author)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] create author", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockAuthor := model.Author{Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			CreateAuthor(mockAuthor).
			Return(model.Author{}, errors.New("ERROR: error while creating author")).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.CreateAuthor(mockAuthor)

		assert.Equal(t, model.Author{}, author)
		assert.NotNil(t, err)
	})
}

func Test_UpdateAuthor(t *testing.T) {

	t.Run("SERVICE [success case] update author", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockAuthor := model.Author{Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			UpdateAuthor(mockID, mockAuthor).
			Return(mockAuthor, nil).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.UpdateAuthor(mockID, mockAuthor)

		assert.NotNil(t, author)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] update author", func(t *testing.T) {
		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockAuthor := model.Author{Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			UpdateAuthor(mockID, mockAuthor).
			Return(model.Author{}, errors.New("ERROR: error while updating author")).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.UpdateAuthor(mockID, mockAuthor)

		assert.Equal(t, model.Author{}, author)
		assert.NotNil(t, err)
	})

}

func Test_DeleteAuthor(t *testing.T) {

	t.Run("SERVICE [success case] delete author", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"
		mockAuthor := model.Author{Name: "author name 1", Surname: "Author surname 1"}

		mockRepository.
			EXPECT().
			DeleteAuthor(mockID).
			Return(mockAuthor, nil).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.DeleteAuthor(mockID)

		assert.NotNil(t, author)
		assert.Nil(t, err)
	})

	t.Run("SERVICE [error case] delete author", func(t *testing.T) {

		mockController := gomock.NewController(t)
		mockRepository := mock.NewMockIRepository(mockController)

		mockID := "1"

		mockRepository.
			EXPECT().
			DeleteAuthor(mockID).
			Return(model.Author{}, errors.New("ERROR: error while deleting author")).
			Times(1)

		service := service.NewService(mockRepository)
		author, err := service.DeleteAuthor(mockID)

		assert.Equal(t, model.Author{}, author)
		assert.NotNil(t, err)
	})

}
