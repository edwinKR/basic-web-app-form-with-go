package server

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func MySQLConnect() (*sql.DB, error) {
	//	Connecting to MySQL db via Google CloudSQL

	var (
		dbUser    = "NEED TO PUT DB_USER_NAME"
		dbPwd     = "NEED TO PUT DB_USER_PWD"
		dbName    = "NEED TO PUT DB_NAME IN CLOUD SQL INSTANCE "
		instanceConnectionName = "NEED TO PUT CONNECTION_NAME OF CLOUD SQL INSTANCE"
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	driverName := "mysql"

	// Use this dataSource URI for App Engine connecting to the Cloud SQL instance.
	dataSource := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	// Use this dataSource URI if connecting from local machine to the Cloud SQL instance. (FYI - Your local machine IP address must be whitelisted on the Cloud SQL instance.)
	//dataSource := "someDB_USER_NAME:someDB_USER_PWD(someCLOUD_SQL_INSTANCE_PUBLIC_IP:3306)/someDB_NAME"

	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		log.Printf("err, %v", err)
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




