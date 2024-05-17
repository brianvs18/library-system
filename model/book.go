package model

type Book struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	AuthorId  string `json:"author_id"`
	Publisher string `json:"publisher"`
	Genre     string `json:"genre"`
	Quantity  int    `json:"quantity"`
}
