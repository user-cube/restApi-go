package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/user-cube/restApi-go/entities"
	"math/rand"
	"net/http"
	"strconv"
)

var Books []entities.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Books)
}
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Books {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book entities.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = rand.Intn(10000000) // Mock id - not safe
	Books = append(Books, book)
	json.NewEncoder(w).Encode(book)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			Books = append(Books[:index], Books[index+1:]...)
			var book entities.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = i
			Books = append(Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		i, _ := strconv.Atoi(params["id"])
		if item.ID == i {
			Books = append(Books[:index], Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}
