package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBook(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	vars := mux.Vars(httpRequest)
	id, _ := strconv.Atoi(vars["id"])
	log.Println(id)
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	for _, b := range Books {
		if b.Id == id {
			json.NewEncoder(httpResponseWriter).Encode(b)
			return
		}
	}
	json.NewEncoder(httpResponseWriter).Encode(nil)
}

func getAllBooks(httpResponseWriter http.ResponseWriter, _ *http.Request) {
	log.Println("Endpoint getAllBooks")
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(httpResponseWriter).Encode(Books)
}

func homepage(httpResponseWriter http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(httpResponseWriter, "Welcome to Homepage, Check ReadMe for endpoints")
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/books/", getAllBooks)
	myRouter.HandleFunc("/book/{id}", getBook)

	fmt.Println("Serving...")
	http.ListenAndServe(":12345", myRouter)
	// databaseConnection()
}
