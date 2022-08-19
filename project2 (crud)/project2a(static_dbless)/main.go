package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)
// static database less struct as dict
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

type Hall struct{
	ID string `json:"id"`
	Movie string `json:"movie"`
	Capacity string `json:"string"`
}
var movies []Movie
var halls []Hall

func getMovies(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	json.NewEncoder(w).Encode(movies) //encode struct to movies
}

func getHalls(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	type Halla struct{
		ID string `json:"id"`
		Movie *Movie `json:"movie"`
		Capacity string `json:"string"`
	}
		var hals []Halla
		for i,j := range halls{
			var hal Halla
			hal.ID=halls[i].ID
			hal.Movie = &Movie{ID:"", Isbn:"", Title:"", Director: &Director{Firstname: "", Lastname:""}}
			for _,l := range movies{
				if(j.Movie== l.Title){
					hal.Movie.ID=l.ID
					hal.Movie.Isbn=l.Isbn
					hal.Movie.Title=l.Title
					hal.Movie.Director.Firstname=l.Director.Firstname
					hal.Movie.Director.Lastname=l.Director.Lastname
				}
			}
			hal.Capacity=halls[i].Capacity
			hals = append(hals, hal)
		}
	
	
	json.NewEncoder(w).Encode(hals) //encode struct to movies
}

func getMoviya(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	type Movieya struct{
		ID string `json:"id"`
		Title string `json:"title"`
	}
	var mov []Movieya
	for i,_ :=range movies{
		var m Movieya
		m.ID=movies[i].ID
		m.Title=movies[i].Title
		mov = append(mov, m)
	}
	json.NewEncoder(w).Encode(mov) //encode struct to movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	params := mux.Vars(r)
	for index, item :=range movies{
		if item.ID == params["id"]{
			movies =append(movies[:index],movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}
}


func getMovieDirector(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	params := mux.Vars(r)
	var dl []Movie
	for _, item :=range movies{
		if item.Director.Firstname == params["name"] && item.Director.Lastname == params["sname"]{
			var movie Movie=item
			dl = append(dl,movie)
		}
	}
	json.NewEncoder(w).Encode(dl)
}

func getMovie(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	params := mux.Vars(r)
	for _, item :=range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	params := mux.Vars(r)
	for index, item :=range movies{
		if item.ID == params["id"]{
			var movie Movie
			_= json.NewDecoder(r.Body).Decode(&movie)
			movie.ID =params["id"]
			movies =append(movies[:index],movies[index+1:]...)
			movies =append(movies, movie)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) { //res, req * is pointer
	w.Header().Set("Content-Type","application/json") //convert struct to appln/json
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID =strconv.Itoa(rand.Intn(100000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
	}


func main() {
	r := mux.NewRouter()  //route to address
	
	//data
	movies = append(movies, Movie{ID:"1", Isbn:"3454455", Title:"Movie one", Director: &Director{Firstname: "Poppy", Lastname:"Singh"}})
	movies = append(movies, Movie{ID:"2", Isbn:"3454455", Title:"Movie two", Director: &Director{Firstname: "Pogli", Lastname:"Aurat"}})
	movies = append(movies, Movie{ID:"3", Isbn:"345423455", Title:"Movie three", Director: &Director{Firstname: "Poppy", Lastname:"Singh"}})
	movies = append(movies, Movie{ID:"4", Isbn:"3451234455", Title:"Movie four", Director: &Director{Firstname: "Pogli", Lastname:"Aurat"}})
	movies = append(movies, Movie{ID:"5", Isbn:"34435454455", Title:"Movie five", Director: &Director{Firstname: "Poppy", Lastname:"Ki Bebo"}})
	movies = append(movies, Movie{ID:"6", Isbn:"7653454455", Title:"Movie six", Director: &Director{Firstname: "Pogli", Lastname:"Aurat"}})
	movies = append(movies, Movie{ID:"7", Isbn:"34435454455", Title:"Movie seven", Director: &Director{Firstname: "Poppy", Lastname:"Ki Bebo"}})
	movies = append(movies, Movie{ID:"8", Isbn:"7653454455", Title:"Movie eight", Director: &Director{Firstname: "Poppy", Lastname:"Kutta"}})

	halls=append(halls, Hall{ID:"1", Movie:"Movie one", Capacity:"10000"})
	halls=append(halls, Hall{ID:"1", Movie:"Movie zero", Capacity:"5000"})
	halls=append(halls, Hall{ID:"1", Movie:"Movie two", Capacity:"15000"})
	halls=append(halls, Hall{ID:"1", Movie:"Movie four", Capacity:"50000"})

	r.HandleFunc("/movies", getMovies).Methods("GET")   // basic routes
	r.HandleFunc("/hall", getHalls).Methods("GET")   // basic routes
	r.HandleFunc("/movieya", getMoviya).Methods("GET")   // basic routes
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")   
	r.HandleFunc("/moviedirector/{name}/{sname}", getMovieDirector).Methods("GET")   
	r.HandleFunc("/movie", createMovie).Methods("POST")   
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")   
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	
	fmt.Println("Start server at port 8069\n")
	log.Fatal(http.ListenAndServe(":8069", r))  //put port
}