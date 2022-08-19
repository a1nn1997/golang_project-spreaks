package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/a1nn1997/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("server is running on port 8069")
	log.Fatal(http.ListenAndServe("localhost:8069", r))
}