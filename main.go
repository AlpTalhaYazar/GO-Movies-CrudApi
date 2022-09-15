package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main() {
	r := mux.NewRouter()

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
