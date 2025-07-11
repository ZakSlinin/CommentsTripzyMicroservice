package main

import (
	"CommentsTripzyMicroservice/internal/services/db"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	dbInstance := initDB()
	db.InitService(dbInstance)

	router := mux.NewRouter()

	//routes

	//runner
	fmt.Println("Server is running on http://localhost:8080") // в продакшене поменять
	http.ListenAndServe(":8080", router)
}

func initDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		fmt.Println("DATABASE_URL environment variable not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error to open DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error to ping DB:", err)
	}

	return db
}
