package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/theblindoracle/aphrodite-backend/internal/database"
	"github.com/theblindoracle/aphrodite-backend/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DB_PATH")
	if dbURL == "" {
		log.Printf("Could not env variable DB_URL")
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Printf("Could not env variable PORT")
		return
	}

	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := server.Config{
		Db: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /notes", cfg.HandlerGetNotes)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
