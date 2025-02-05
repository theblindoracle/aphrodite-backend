package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/theblindoracle/aphrodite-backend/internal/database"
)

const dbFile string = "app.db"

type Config struct {
	db *database.Queries
}

func main() {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("could not open database: %v", err)
	}

	dbQueries := database.New(db)

	cfg := Config{
		db: dbQueries,
	}

	_, err = cfg.db.CreateNote(context.Background(), "I Love you!")
	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}
	_, err = cfg.db.CreateNote(context.Background(), "You're the best")
	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}

	notes, err := cfg.db.GetAllNotes(context.Background())

	for idx, note := range notes {

		log.Printf("Note %v", idx)
		log.Printf("*ID			%v", note.ID)
		log.Printf("*Created At %v", note.CreatedAt)
		log.Printf("*Updated At %v", note.UpdatedAt)
		log.Printf("*Note		%v", note.Note)
	}
}
