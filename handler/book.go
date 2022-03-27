package handler

import (
	"encoding/json"
	"net/http"

	"github.com/parvineyvazov/book-api/model"

	"github.com/gorilla/mux"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	books, err := h.service.GetBooks()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	book, err := h.service.GetBook(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func (h *Handler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	books, err := h.service.SearchBooks(vars["search_text"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, books)
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {

	var book model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	created_book, err := h.service.CreateBook(book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, created_book)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var book model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	updated_book, err := h.service.UpdateBook(vars["id"], book)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updated_book)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	deleted_book, err := h.service.DeleteBook(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, deleted_book)
}
