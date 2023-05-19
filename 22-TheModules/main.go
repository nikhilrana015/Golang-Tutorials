package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeRouter).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func HomeRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang world</h1>"))
}
