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

func Test_GetAuthors(t *testing.T) {

	t.Run("HANDLER [success case] get authors", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Authors{
			{Id: "1", Name: "author name 1", Surname: "Author surname 1"},
			{Id: "2", Name: "author name 2", Surname: "Author surname 2"},
		}

		mockService.
			EXPECT().
			GetAuthors().
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/authors", http.NoBody)
		res := httptest.NewRecorder()

		handler.GetAuthors(res, req)

		var expectedResBody model.Authors
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData[0].Id, expectedResBody[0].Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] get authors", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockService.
			EXPECT().
			GetAuthors().
			Return(nil, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/authors", http.NoBody)
		res := httptest.NewRecorder()

		handler.GetAuthors(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_GeAuthor(t *testing.T) {

	t.Run("HANDLER [success case] get author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Author{Id: "2", Name: "author name 2", Surname: "Author surname 2"}

		mockService.
			EXPECT().
			GetAuthor(mockID).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		res := httptest.NewRecorder()

		handler.GetAuthor(res, req)

		var expectedResBody model.Author
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] get author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"

		mockService.
			EXPECT().
			GetAuthor(mockID).
			Return(model.Author{}, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/Author", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})

		res := httptest.NewRecorder()

		handler.GetAuthor(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_SearchAuthors(t *testing.T) {

	t.Run("HANDLER [success case] search authors", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockSearchText := "tit"
		mockReturnData := model.Authors{
			{Id: "1", Name: "author name 1", Surname: "Author surname 1"},
			{Id: "2", Name: "author name 2", Surname: "Author surname 2"},
		}

		mockService.
			EXPECT().
			SearchAuthors(mockSearchText).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author/search", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"search_text": mockSearchText,
		})

		res := httptest.NewRecorder()

		handler.SearchAuthors(res, req)

		var expectedResBody model.Authors
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData[0].Id, expectedResBody[0].Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] search authors", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockSearchText := "tit"

		mockService.
			EXPECT().
			SearchAuthors(mockSearchText).
			Return(nil, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author/search", http.NoBody)

		req = mux.SetURLVars(req, map[string]string{
			"search_text": mockSearchText,
		})

		res := httptest.NewRecorder()

		handler.SearchAuthors(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_CreateAuthor(t *testing.T) {

	t.Run("HANDLER [success case] create author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockService.
			EXPECT().
			CreateAuthor(mockReturnData).
			Return(mockReturnData, nil).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", &buf)
		res := httptest.NewRecorder()

		handler.CreateAuthor(res, req)

		var expectedResBody model.Author
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] create author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockReturnData := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockService.
			EXPECT().
			CreateAuthor(mockReturnData).
			Return(model.Author{}, errors.New("ERROR: error")).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", &buf)
		res := httptest.NewRecorder()

		handler.CreateAuthor(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_UpdateAuthor(t *testing.T) {

	t.Run("HANDLER [success case] update author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockService.
			EXPECT().
			UpdateAuthor(mockID, mockReturnData).
			Return(mockReturnData, nil).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", &buf)

		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})

		res := httptest.NewRecorder()

		handler.UpdateAuthor(res, req)

		var expectedResBody model.Author
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] update author", func(t *testing.T) {
		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockService.
			EXPECT().
			UpdateAuthor(mockID, mockReturnData).
			Return(model.Author{}, errors.New("ERROR: error")).
			Times(1)

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(mockReturnData)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", &buf)

		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})

		res := httptest.NewRecorder()

		handler.UpdateAuthor(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}

func Test_DeleteAuthor(t *testing.T) {

	t.Run("HANDLER [success case] delete author", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"
		mockReturnData := model.Author{Id: "1", Name: "author name 1", Surname: "Author surname 1"}

		mockService.
			EXPECT().
			DeleteAuthor(mockID).
			Return(mockReturnData, nil).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", http.NoBody)
		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})
		res := httptest.NewRecorder()

		handler.DeleteAuthor(res, req)

		var expectedResBody model.Author
		err := json.Unmarshal(res.Body.Bytes(), &expectedResBody)

		assert.Nil(t, err, "Error on json unmarshal")
		assert.Equal(t, mockReturnData.Id, expectedResBody.Id)
		assert.Equal(t, http.StatusOK, res.Result().StatusCode)
	})

	t.Run("HANDLER [error case] delete author", func(t *testing.T) {

		mockService := mock.NewMockIService(gomock.NewController(t))

		mockID := "1"

		mockService.
			EXPECT().
			DeleteAuthor(mockID).
			Return(model.Author{}, errors.New("ERROR: error")).
			Times(1)

		handler := handler.NewHandler(mockService)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/author", http.NoBody)
		req = mux.SetURLVars(req, map[string]string{
			"id": mockID,
		})
		res := httptest.NewRecorder()

		handler.DeleteAuthor(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
	})

}
