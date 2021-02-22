package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func getBooks(w http.ResponseWriter, r *http.Request)  {}
func getBook(w http.ResponseWriter, r *http.Request)  {}
func createBook(w http.ResponseWriter, r *http.Request)  {}
func updateBook(w http.ResponseWriter, r *http.Request)  {}
func deleteBook(w http.ResponseWriter, r *http.Request)  {}

func handleRequests()  {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/books", getBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", getBook).Methods("GET")
	myRouter.HandleFunc("/books", createBook).Methods("POST")
	myRouter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
func main(){
	handleRequests()
}
