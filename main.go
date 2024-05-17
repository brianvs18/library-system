package main

import (
	"github.com/gorilla/mux"
	"library-system/config"
	"library-system/controller"
	"log"
	"net/http"
)

func main() {
	err := config.Connection()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/authors", controller.GetAllAuthors).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

}
