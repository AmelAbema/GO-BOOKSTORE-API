package routes

import (
	"GO-BOOKSTORE-API/src/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.createBook).Methods("POST")
	router.HandleFunc("/books", controllers.getAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookid}", controllers.getBookById).Methods("GET")
	router.HandleFunc("/books/{bookid}", controllers.updateBook).Methods("PUT")
	router.HandleFunc("/books/{bookid}", controllers.deleteBook).Methods("DELETE")
}
