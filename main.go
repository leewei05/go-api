package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func initDB() {
	_ = godotenv.Load("config.env")

	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPwd := os.Getenv("PG_PWD")
	dbName := os.Getenv("PG_DB")

	pgStr := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPwd, dbName)
	fmt.Println(pgStr)

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPwd == "" || dbName == "" {
		log.Panicf("Missing config parameters: %v", pgStr)
	}

	_, err := sql.Open("postgres", pgStr)
	if err != nil {
		log.Panic("Cannot open PostgreSQL database")
	}
}

func main() {
	initDB()

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/v1/", get).Methods("GET")

	http.Handle("/", r)

	serverPort := os.Getenv("HTTP_PORT")
	if serverPort == "" {
		log.Panic("Null HTTP port value")
	}

	port := fmt.Sprintf(":%v", serverPort)

	s := &http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("HTTP server running on port %v", port)
	log.Fatal(s.ListenAndServe())
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
