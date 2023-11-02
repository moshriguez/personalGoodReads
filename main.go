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
	Title      string    `bson:"title" json:"title"`
	Author     string    `bson:"author" json:"author"`
	Read       bool      `bson:"read" json:"read"`
	ReadOn     time.Time `bson:"read_on" json:"read_on"`
	LostInMove bool      `bson:"lost_in_move" json:"lost_in_move"`
}

var library []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetBooks: not implemented yet!")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AddBook: not implemented yet!")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdateBook: not implemented yet!")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeleteBook: not implemented yet!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GetBook: not implemented yet!")
}

func main() {
	// init router
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my library.")
	}).Methods("GET")
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books", AddBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
