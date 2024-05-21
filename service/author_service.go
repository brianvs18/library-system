package service

import (
	"library-system/config"
	"library-system/model"
	"library-system/util"
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

func SaveAuthor(author model.Author) error {
	query := "INSERT INTO Authors(Id,Name) VALUES(?,?)"

	_, err := config.DB.Exec(query, util.GenerateId(), author.Name)
	if err != nil {
		return err
	}
	return nil
}
