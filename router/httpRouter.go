package router

import (
	"github.com/gorilla/mux"
	"github.com/user-cube/restApi-go/controllers"
	"github.com/user-cube/restApi-go/entities"
	"log"
	"net/http"
)

func HandleRequests() {
	controllers.Books = append(controllers.Books, entities.Book{Title: "Nice book", ISBN: "133NFWEF2223E23", ID: 1, Author: &entities.Author{FirstName: "Miguel", LastName: "Galhetas"}})
	controllers.Books = append(controllers.Books, entities.Book{Title: "Nice book 2", ISBN: "133NFWEF2223323", ID: 2 , Author: &entities.Author{FirstName: "Miguel", LastName: "Galhetas"}})

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	myRouter.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	myRouter.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	myRouter.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	myRouter.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
