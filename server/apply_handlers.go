package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Applicant struct {
	Customerid string `json:"customerId"`
	LName		string `json:"lname"`
	FName		string	`json:"fname"`
	Contact		string `json:"contact"`
}

var db *sql.DB
var errInit error

// Initiating DB instance
func InitDB() {
	fmt.Println("Initiating DB!!!!")
	db, errInit = MySQLConnect()
	if errInit != nil {
		fmt.Println("Error in InitDB...")
		fmt.Println(errInit)
	}
	//defer db.Close()

	fmt.Println("DB>>>>", db)
}

// GET handler serving all existing applicants.
func GetApplicantsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Request /status")
	rows, err := db.Query(`SELECT * FROM candidate;`)
	if err != nil {
		fmt.Println("Query error...\n")
		fmt.Errorf("%v", err)
	}

	//defer db.Close()

	//Testing================Start=============================
	//fmt.Println("###########TESTING STARTS~~~~~~~~")
	////Data to be used in query
	//var (
	//	customer_id string
	//	lname string
	//	fname string
	//	contact string
	//)
	//
	//for rows.Next() {
	//	err = rows.Scan(&customer_id, &lname, &fname, &contact)
	//	if err != nil {
	//		fmt.Println("Uh oh....", err)
	//	}
	//	fmt.Println("***Retreived data**** \n", customer_id, lname, fname, contact)
	//}
	//fmt.Println("Testing done~~~~~~~~")
	//Testing===========END==================================


	defer rows.Close()
	var applicants []Applicant

	//query
	for rows.Next() {
		// For each row returned by the table, create a pointer to the single applicant.
		applicant := &Applicant{}

		err = rows.Scan(&applicant.Customerid, &applicant.LName, &applicant.FName, &applicant.Contact)
		if err != nil {
			fmt.Println("Uh oh....", err)
		}
		fmt.Println("***Retreived data**** \n")
		applicants = append(applicants, *applicant)
		fmt.Printf("%#v >>>> \n", applicants)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("rows Error...", err)
	}
	fmt.Println("Done!!")


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

	stmt, err2 := db.Prepare(`INSERT INTO candidate(lname, fname, contact) VALUES (?,?,?)`)
	if err2 != nil {
		fmt.Println("Post query err...\n")
		fmt.Println(err2)
	}
	stmt.Exec(newApplicant.LName, newApplicant.FName, newApplicant.Contact)
	fmt.Println("After INSERT INTO statement ====> \n", newApplicant)

	// Redirecting the user to the index.html page.
	http.Redirect(w, r, "/", http.StatusFound)
}

// Update handler to edit applicant information.
func UpdateApplicantHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE request /edit")

	selectedApplicant := Applicant{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	selectedApplicant.Customerid = r.Form.Get("customerId")
	selectedApplicant.LName = r.Form.Get("lname")
	selectedApplicant.FName = r.Form.Get("fname")
	selectedApplicant.Contact = r.Form.Get("contact")
	fmt.Printf("%#v $$$$$$$$$>>> ", selectedApplicant)
	stmt, err := db.Prepare(`UPDATE candidate SET lname=?, fname=?, contact=? WHERE customer_id=?`)
	if err != nil {
		fmt.Println("Update STMT error...", err)
	}
	stmt.Exec(selectedApplicant.LName, selectedApplicant.FName, selectedApplicant.Contact, selectedApplicant.Customerid)
	defer stmt.Close()
	http.Redirect(w, r, "/", http.StatusFound)
}

// Delete handler to delete applicant.
func DeleteApplicantHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE request /delete")

	type incomingBody struct {
		Uid string `json:uid"`
	}
	var reqBody incomingBody

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reqBody)
	fmt.Printf("INCOMING!!!!!!!==== \n %#v", reqBody.Uid)

	stmt, err := db.Prepare("DELETE FROM candidate WHERE customer_id=?")
	if err != nil {
		fmt.Println("Delete STMT error...", err)
	}
	stmt.Exec(reqBody.Uid)
	defer stmt.Close()
	http.Redirect(w, r, "/", http.StatusFound)
}