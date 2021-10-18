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

type Movie struct {
	ID string`json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies [] Movie

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1",Isbn:"438227",Title:"Movie One",Director:&Director{Firstname: "john",Lastname: "Doe"}})
	r.handleFunc("movies/",getMovies).Methods("GET")
	r.handleFunc("movies/{id}",getMovie).Metohds("GET")
	r.handleFunc("/movie",createMovie).Methods("POST")
	r.handleFunc("/movies/{id}",updateMovies).Methods("PUT")
	r.handleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Print("start server at port 8000\n")
	log.Fatal(http.listenAndServer(":8000",r))
}
