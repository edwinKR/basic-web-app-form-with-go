package main

import (
	"basic-web-app-form-with-go/server"
	//"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Creating all routes
func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/apply", server.CreateNewApplicantHandler).Methods("POST")
	r.HandleFunc("/status", server.GetApplicantsHandler).Methods("GET")
	r.HandleFunc("/edit", server.UpdateApplicantHandler).Methods("POST")
	r.HandleFunc("/delete", server.DeleteApplicantHandler).Methods("DELETE")
	staticAssetsDirectory := http.Dir("./client")
	//staticFileHandler := http.FileServer(staticAssetsDirectory)
	staticFileHandler := http.StripPrefix("", http.FileServer(staticAssetsDirectory))
	r.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	//r.Handle("/", staticFileHandler).Methods("GET")

	return r
}

// main() function - This is the entry point of my executable program. This is where it starts the application.
func main() {
	fmt.Println("---------DB Server running???--------")
	//MyDB, Err := server.MySQLConnect()
	//if Err != nil {
	//	fmt.Println(Err)
	//}
	//fmt.Println("DB>>>>", MyDB)

	server.InitDB()

	router := newRouter()
	fmt.Println("Serving on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}