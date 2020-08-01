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
	r.HandleFunc("/", server.HomePageHandler).Methods("GET")

	staticAssetsDirectory := http.Dir("./client/")
	staticFileHandler := http.StripPrefix("/client/", http.FileServer(staticAssetsDirectory))
	r.PathPrefix("/client/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/status", server.GetApplicantsHandler).Methods("GET")
	r.HandleFunc("/apply", server.CreateNewApplicantHandler).Methods("POST")

	return r
}

func main() {
	fmt.Println("==== Server Running ====")

	r := newRouter()
	fmt.Println("Serving on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}