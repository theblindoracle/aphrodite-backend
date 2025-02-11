package config

import "github.com/theblindoracle/aphrodite-backend/internal/database"

type Config struct {
	Db *database.Queries
}
