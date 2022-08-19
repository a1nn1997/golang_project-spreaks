package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/a1nn1997/go-bookstore/pkg/utils"
	"github.com/a1nn1997/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBook()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err:= strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db:=models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLibrary(w http.ResponseWriter, r *http.Request){
	newLibrarys:=models.GetAllLibrary()
	res, _ :=json.Marshal(newLibrarys)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetLibraryById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	LibraryId := vars["LibraryId"]
	ID, err:= strconv.ParseInt(LibraryId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	LibraryDetails, _:= models.GetLibraryById(ID)
	res, _ := json.Marshal(LibraryDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateLibrary(w http.ResponseWriter, r *http.Request){
	CreateLibrary := &models.Library{}
	utils.ParseBody(r, CreateLibrary)
	b:= CreateLibrary.CreateLibrary()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteLibrary(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	LibraryId := vars["LibraryId"]
	ID, err := strconv.ParseInt(LibraryId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	Library := models.DeleteLibrary(ID)
	res, _ := json.Marshal(Library)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLibrary(w http.ResponseWriter, r *http.Request){
	var updateLibrary = &models.Library{}
	utils.ParseBody(r, updateLibrary)
	vars := mux.Vars(r)
	LibraryId := vars["LibraryId"]
	ID, err := strconv.ParseInt(LibraryId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	LibraryDetails, db:=models.GetLibraryById(ID)
	if updateLibrary.Name != ""{
		LibraryDetails.Name = updateLibrary.Name
	}
	if updateLibrary.Book != ""{
		LibraryDetails.Book = updateLibrary.Book
	}
	if updateLibrary.Capacity != ""{
		LibraryDetails.Capacity = updateLibrary.Capacity
	}
	db.Save(&LibraryDetails)
	res, _ := json.Marshal(LibraryDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//NewBook me model Book  //json.Marshal covert struct to json
// contentType was changed