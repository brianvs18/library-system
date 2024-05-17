package controller

import (
	"encoding/json"
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
