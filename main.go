package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"./server"
)

// Creating all routes
func newRouter() *mux.Router {
	r := mux.NewRouter()

	staticAssetsDirectory := http.Dir("./client/")
	staticFileHandler := http.FileServer(staticAssetsDirectory)
	// staticFileHandler := http.StripPrefix("/client/", http.FileServer(staticAssetsDirectory))
	// r.PathPrefix("/client/").Handler(staticFileHandler).Methods("GET")

	r.Handle("/", staticFileHandler).Methods("GET")

	r.HandleFunc("/status", server.GetApplicantsHandler).Methods("GET")
	r.HandleFunc("/apply", server.CreateNewApplicantHandler).Methods("POST")

	return r
}

// main() function - This is the entry point of my executable program. This is where it starts the application.
func main() {
	fmt.Println("==== Server Running ====")

	router := newRouter()
	fmt.Println("Serving on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}