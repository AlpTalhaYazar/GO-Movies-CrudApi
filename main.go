package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
