package main

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetalpbalkan/go-linq"
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
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	selectedMovie := linq.From(movies).FirstWithT(func(x Movie) bool {
		return x.ID == params["id"]
	})

	json.NewEncoder(w).Encode(selectedMovie)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie Movie

	json.NewDecoder(r.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(1000000))

	movies = append(movies, newMovie)

	json.NewEncoder(w).Encode(newMovie)
}
func main() {
	r := mux.NewRouter()

	movies = append(movies,
		Movie{
			ID:    "1",
			Isbn:  "54215",
			Title: "Psycho",
			Director: &Director{
				FirstName: "Alfred",
				LastName:  "Hitchcock"}})

	movies = append(movies,
		Movie{
			ID:    "2",
			Isbn:  "62622",
			Title: "2001: A Space Odyssey",
			Director: &Director{
				FirstName: "Stanley",
				LastName:  "Kubrick"}})

	movies = append(movies,
		Movie{
			ID:    "3",
			Isbn:  "816692",
			Title: "Interstellar",
			Director: &Director{
				FirstName: "Christopher",
				LastName:  "Nolan"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
