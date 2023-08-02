package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBDriver   = "mysql"
	DBUser     = "your_database_username"
	DBPassword = "your_database_password"
	DBName     = "your_database_name"
)

var db *sql.DB

func main() {
	// Open database connection
	dbConn, err := sql.Open(DBDriver, fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, DBName))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer dbConn.Close()

	// Test the database connection
	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	db = dbConn

	// Start the server
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/read", ReadHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/delete", DeleteHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
