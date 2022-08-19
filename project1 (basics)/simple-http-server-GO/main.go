package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {   //post req able to parse form
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")  //parse value from form.html
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	if r.URL.Path != "/hello" {  //path error dont match
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {  //mismatch error get 
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //FileServer checkout static file
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8069\n")
	if err := http.ListenAndServe(":8069", nil); err != nil {
		log.Fatal(err)
	}//connect to server for with 8069
}
