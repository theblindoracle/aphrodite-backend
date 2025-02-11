package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/theblindoracle/aphrodite-backend/internal/config"
	"github.com/theblindoracle/aphrodite-backend/internal/database"
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
	db, err := sql.Open("sqlite3", dbURL)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := config.Config{
		Db: dbQueries,
	}

	_, err = cfg.Db.CreateNote(context.Background(), "I Love you!")
	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}
	_, err = cfg.Db.CreateNote(context.Background(), "You're the best")
	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}

	notes, err := cfg.Db.GetAllNotes(context.Background())

	for idx, note := range notes {

		log.Printf("Note %v", idx)
		log.Printf("*ID			%v", note.ID)
		log.Printf("*Created At %v", note.CreatedAt)
		log.Printf("*Updated At %v", note.UpdatedAt)
		log.Printf("*Note		%v", note.Note)
	}
}
