package routes

import (
	"github.com/gorilla/mux"
	"github.com/a1nn1997/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/library/", controllers.CreateLibrary).Methods("POST")
	router.HandleFunc("/library/", controllers.GetLibrary).Methods("GET")
	router.HandleFunc("/library/{LibraryId}", controllers.GetLibraryById).Methods("GET")
	router.HandleFunc("/library/{LibraryId}", controllers.UpdateLibrary).Methods("PUT")
	router.HandleFunc("/library/{LibraryId}", controllers.DeleteLibrary).Methods("DELETE")
}