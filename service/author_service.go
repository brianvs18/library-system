package service

import (
	"library-system/config"
	"library-system/model"
)

func GetAllAuthors() ([]model.Author, error) {
	query := "SELECT * FROM Authors"

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []model.Author
	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}
