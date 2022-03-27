package server

import (
	"fmt"
	"net/http"

	"github.com/parvineyvazov/book-api/handler"
	"github.com/parvineyvazov/book-api/repository"
	"github.com/parvineyvazov/book-api/service"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var LOCAL_STORAGE_API_URL string = "/api/v1"
var DB_API_URL string = "/api/v2"

func setupLocalStorageRoutes(myRouter *mux.Router) {
	repository := repository.NewLocalRepository()
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	// BOOK endpoints
	myRouter.HandleFunc(localUrl("books"), handler.GetBooks).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("book/{id}"), handler.GetBook).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("book/search/{search_text}"), handler.SearchBooks).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("book"), handler.CreateBook).Methods(http.MethodPost)
	myRouter.HandleFunc(localUrl("book/{id}"), handler.UpdateBook).Methods(http.MethodPut)
	myRouter.HandleFunc(localUrl("book/{id}"), handler.DeleteBook).Methods(http.MethodDelete)

	// AUTHOR endpoints
	myRouter.HandleFunc(localUrl("authors"), handler.GetAuthors).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("author/{id}"), handler.GetAuthor).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("author/search/{search_text}"), handler.SearchAuthors).Methods(http.MethodGet)
	myRouter.HandleFunc(localUrl("author"), handler.CreateAuthor).Methods(http.MethodPost)
	myRouter.HandleFunc(localUrl("author/{id}"), handler.UpdateAuthor).Methods(http.MethodPut)
	myRouter.HandleFunc(localUrl("author/{id}"), handler.DeleteAuthor).Methods(http.MethodDelete)
}

func (s *Server) StartServer(port int) error {

	myRouter := mux.NewRouter().StrictSlash(true)

	setupLocalStorageRoutes(myRouter)

	httpHandler := cors.Default().Handler(myRouter)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), httpHandler)

	return err
}

func localUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", LOCAL_STORAGE_API_URL, endpoint)
}

func dbUrl(endpoint string) string {
	return fmt.Sprintf("%s/%s", DB_API_URL, endpoint)
}
