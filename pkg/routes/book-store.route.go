package routes

import (
	"github.com/abdelhalim97/CRUD/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.Books).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.BookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("Delete")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PATCH")
}
