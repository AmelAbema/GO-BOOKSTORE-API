package routes

import (
	"GO-BOOKSTORE-API/src/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
