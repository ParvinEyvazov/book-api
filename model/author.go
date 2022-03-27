package model

type Author struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
	Books   Books  `json:"books,omitempty"`
	// others
}

type Authors []Author
