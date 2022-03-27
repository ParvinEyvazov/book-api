package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/parvineyvazov/book-api/model"
)

// GetAuthors get all authors from DB
func (*RepositoryDB) GetAuthors() (model.Authors, error) {

	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM Author")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var authors []model.Author
	for results.Next() {
		fmt.Println("next")
		var author model.Author
		// for each row, scan the result into our tag composite object
		err = results.Scan(&author.Id, &author.Name, &author.Surname)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		authors = append(authors, author)
	}

	for _, v := range authors {
		fmt.Println("author ", v)
	}

	return authors, nil
}

// GetAuthor get author by id from DB
func (*RepositoryDB) GetAuthor(id string) (model.Author, error) {

	return model.Author{}, nil
}

// SearchAuthor get authors with related search text from DB
func (*RepositoryDB) SearchAuthors(search_text string) (model.Authors, error) {

	return nil, nil
}

// CreateAuthor create author in DB
func (*RepositoryDB) CreateAuthor(author model.Author) (model.Author, error) {

	return model.Author{}, nil
}

// UpdateAuthor updates author in DB
func (*RepositoryDB) UpdateAuthor(id string, author model.Author) (model.Author, error) {

	return model.Author{}, nil
}

// DeleteAuthor deletes author by id from DB
func (*RepositoryDB) DeleteAuthor(id string) (model.Author, error) {

	return model.Author{}, nil
}
