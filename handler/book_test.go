package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/parvineyvazov/book-api/handler"
	"github.com/parvineyvazov/book-api/mock"
	"github.com/parvineyvazov/book-api/model"
	"github.com/stretchr/testify/assert"
)

func Test_GetBooks(t *testing.T) {

	t.Run("HANDLER [success case] get books", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Books{
			{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}},
			{Id: "2", Title: "title 2", Description: "desc 2", Publication_date: "2022", AuthorIDs: []string{"2"}},
		}

		mockService.
			EXPECT().
			GetBooks().
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/books", http.NoBody)
		res := httptest.NewRecorder()

		handler.GetBooks(res, req)

		var expectedResBody model.Books
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData[0].Id, expectedResBody[0].Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] get books", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockService.
			EXPECT().
			GetBooks().
			Return(nil, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/books", http.NoBody)
		res := httptest.NewRecorder()

		handler.GetBooks(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_GetBook(t *testing.T) {

	t.Run("HANDLER [success case] get book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			GetBook(mockID).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		res := httptest.NewRecorder()

		handler.GetBook(res, req)

		var expectedResBody model.Book
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] get book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"

		mockService.
			EXPECT().
			GetBook(mockID).
			Return(model.Book{}, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		res := httptest.NewRecorder()

		handler.GetBook(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_SearchBooks(t *testing.T) {

	t.Run("HANDLER [success case] search books", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockSearchText := "tit"
		mockReturnData := model.Books{
			{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}},
			{Id: "2", Title: "title 2", Description: "desc 2", Publication_date: "2022", AuthorIDs: []string{"2"}},
		}

		mockService.
			EXPECT().
			SearchBooks(mockSearchText).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book/search", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"search_text": mockSearchText,
		})

		res := httptest.NewRecorder()

		handler.SearchBooks(res, req)

		var expectedResBody model.Books
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData[0].Id, expectedResBody[0].Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] search books", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockSearchText := "tit"

		mockService.
			EXPECT().
			SearchBooks(mockSearchText).
			Return(nil, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book/search", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"search_text": mockSearchText,
		})

		res := httptest.NewRecorder()

		handler.SearchBooks(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_CreateBook(t *testing.T) {

	t.Run("HANDLER [success case] create book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			CreateBook(mockReturnData).
			Return(mockReturnData, nil).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", &buf)
		res := httptest.NewRecorder()

		handler.CreateBook(res, req)

		var expectedResBody model.Book
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] create book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			CreateBook(mockReturnData).
			Return(model.Book{}, errors.New("ERROR: error")).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", &buf)
		res := httptest.NewRecorder()

		handler.CreateBook(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_UpdateBook(t *testing.T) {

	t.Run("HANDLER [success case] update book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			UpdateBook(mockID, mockReturnData).
			Return(mockReturnData, nil).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", &buf)

		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})

		res := httptest.NewRecorder()

		handler.UpdateBook(res, req)

		var expectedResBody model.Book
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] update book", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			UpdateBook(mockID, mockReturnData).
			Return(model.Book{}, errors.New("ERROR: error")).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", &buf)

		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})

		res := httptest.NewRecorder()

		handler.UpdateBook(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_DeleteBook(t *testing.T) {

	t.Run("HANDLER [success case] delete book", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Book{Id: "1", Title: "title 1", Description: "desc 1", Publication_date: "2021", AuthorIDs: []string{"1"}}

		mockService.
			EXPECT().
			DeleteBook(mockID).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", http.NoBody)
		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})
		res := httptest.NewRecorder()

		handler.DeleteBook(res, req)

		var expectedResBody model.Book
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] delete book", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"

		mockService.
			EXPECT().
			DeleteBook(mockID).
			Return(model.Book{}, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/book", http.NoBody)
		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})
		res := httptest.NewRecorder()

		handler.DeleteBook(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}
