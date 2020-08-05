package server

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	//"gocloud.dev/mysql"
	//_ "gocloud.dev/mysql/gcpmysql"
)

func MySQLConnect() (*sql.DB, error) {
	//	Connecting to MySQL db via Google CloudSQL
	driverName := "mysql"
	dataSource := "NEED THE RIGHT CONNECTION URI HERE!!!!"
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Printf("err, %v", err)
		//fmt.Errorf("%v", err)
		return nil, err
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping error..\n", err)
		log.Printf("err, %v", err)
		fmt.Errorf("%v", err)
	}
	fmt.Println("---------PING successful! DB Server running!!!!!--------")

	//Testing================Start=============================
	//fmt.Println("###########TESTING STARTS~~~~~~~~")
	//rows, err := db.Query(`SELECT * FROM candidate;`)
	//if err != nil {
	//	fmt.Println("Query error...\n")
	//	fmt.Errorf("%v", err)
	//}
	//
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

	return db, nil
}

