package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getBooks(w http.ResponseWriter, r *http.Request)  {}

func handleRequests()  {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/books", getBooks).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
func main(){
	handleRequests()
}
