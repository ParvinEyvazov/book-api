package model

type Book struct {
	Id               string   `json:"id,omitempty"`
	Title            string   `json:"title,omitempty"`
	Description      string   `json:"description,omitempty"`
	Publication_date string   `json:"publication_date,omitempty"`
	AuthorIDs        []string `json:"authorIDs,omitempty"`
	Authors          Authors  `json:"authors,omitempty"`
	// other
}

type Books []Book
