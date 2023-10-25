package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Read       bool
	ReadOn     time.Time
	LostInMove bool
}

var library []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet!")
}

func main() {
	// init router
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books", AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
