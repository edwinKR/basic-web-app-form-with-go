package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Applicant struct {
	LName		string `json:"lname"`
	FName		string	`json:"fname"`
	Contact		string `json:"contact"`
}

var applicants []Applicant

// GET handler serving all existing applicants.
func GetApplicantsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Request /status")
	applicantsBytes, err := json.Marshal(applicants)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err2 := w.Write(applicantsBytes)
	if err2 != nil {
		fmt.Println(fmt.Errorf("Error: %v", err2))
	}
}

// POST handler to create a new applicant.
func CreateNewApplicantHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST request /apply")
	newApplicant := Applicant{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newApplicant.LName = r.Form.Get("lname")
	newApplicant.FName = r.Form.Get("fname")
	newApplicant.Contact = r.Form.Get("contact")

	applicants = append(applicants, newApplicant)

	// Redirecting the user to the index.html page.
	http.Redirect(w, r, "/", http.StatusFound)
}

// // GET Home Page handler
// func HomePageHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to IFM")
// }
