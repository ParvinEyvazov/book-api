package handler

import (
	"encoding/json"
	"net/http"

	"github.com/parvineyvazov/book-api/model"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {

	authors, err := h.service.GetAuthors()

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, authors)

}

func (h *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	author, err := h.service.GetAuthor(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, author)

}

func (h *Handler) SearchAuthors(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	authors, err := h.service.SearchAuthors(vars["search_text"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, authors)
}

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var author model.Author
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	created_author, err := h.service.CreateAuthor(author)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, created_author)
}

func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var author model.Author
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&author); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	updated_author, err := h.service.UpdateAuthor(vars["id"], author)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updated_author)

}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	deleted_author, err := h.service.DeleteAuthor(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, deleted_author)

}
