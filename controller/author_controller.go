package controller

import (
	"encoding/json"
	"library-system/model"
	"library-system/service"
	"net/http"
)

func GetAllAuthors(w http.ResponseWriter, _ *http.Request) {
	authors, err := service.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonParse, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonParse)
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.SaveAuthor(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
