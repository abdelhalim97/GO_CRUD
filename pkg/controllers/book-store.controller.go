package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdelhalim97/CRUD/pkg/models"
	"github.com/abdelhalim97/CRUD/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func Books(w http.ResponseWriter, r *http.Request) {
	newBooks := models.Books()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func BookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing book")
	}
	bookDetails, _ := models.BookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	BookData := &models.Book()
	utils.ParsedBody(r, CreateBook)
	b := BookData.CreateBook()
	res, _ := json.Marshal(b)
}
