package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
	// director has type pointer to struct type director
	// from where it will get it's data

}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `jason:"lastname"`
}

// Now creating slice of type Movies
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
	// this line is encoding all data from slice movies
	// into json and then sending it to user requesting
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID == param["id"] {
			// when we the entered id from web and it matched with item.id then we will delete it
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// now updating json data
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	param := mux.Vars(r)
	// this will extract variable from reauest that we need
	for _, item := range movies {
		if item.ID == param["id"] {
			// now we have to return it
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	// Now we add this to movies slice
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			// here first deleting that index and then we are adding new
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000000))
			// Now we add this to movies slice
			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()
	// Now appending data inside the movies
	// as we know it has type Movies struct

	movies = append(movies, Movie{ID: "1", Isbn: "43877", Title: "Movie one", Director: &Director{FirstName: "john",
		LastName: "Doe"}})
	// here we are using &director because it will be sending address of this object to where it is needed
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "34856",
		Title: "In bruges",
		Director: &Director{
			FirstName: "Guy",
			LastName:  "Ritiche",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("stating server at port 8080")
	log.Fatal(http.ListenAndServe(":8000", r))
	// this is how we start server in golang
	// using http.ListenandServe function.

}
