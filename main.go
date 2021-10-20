package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResposeWriter, r *http.Request) {
	w.Header().set("Context-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteMovie(w http.ResposeWriter, r *http.Request) {
	w.Header().set("Context-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Context type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().set("Context type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	//初期値の設定
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie One", Director: &Director{Firstname: "terry", Lastname: "watson"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45678", Title: "Movie Two", Director: &Director{Firstname: "billy", Lastname: "john"}})
	r.handleFunc("movies/", getMovies).Methods("GET")
	r.handleFunc("movies/{id}", getMovie).Metohds("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("start server at port 8000\n")
	log.Fatal(http.listenAndServer(":8000", r))
}
