package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB

	serverPort = os.Getenv("HTTP_PORT")
	dbHost     = os.Getenv("PG_HOST")
	dbPort     = os.Getenv("PG_PORT")
	dbUser     = os.Getenv("PG_USER")
	dbPwd      = os.Getenv("PG_PWD")
	dbName     = os.Getenv("PG_DB")
)

func initDB() {
	pgStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName)

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPwd == "" || dbName == "" {
		log.Panicf("Missing config parameters: %v", pgStr)
	}

	db, err := sql.Open("postgres", pgStr)
	if err != nil {
		log.Panic("Cannot open PostgreSQL database")
	}
	defer db.Close()
}

func main() {
	initDB()

	if serverPort == "" {
		log.Panic("Null HTTP port value")
	}

	port := fmt.Sprintf(":%v", serverPort)
	log.Printf("HTTP server running on port %v", port)
	http.ListenAndServe(port, nil)
}
