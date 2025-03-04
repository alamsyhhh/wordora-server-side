package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("⚠️ .env file not found, using environment variables")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, logError("DATABASE_URL is not set in the environment")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, logError("Failed to connect to the database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, logError("Database connection failed: %v", err)
	}

	log.Println("✅ Successfully connected to the database!")
	return db, nil
}

func logError(format string, v ...interface{}) error {
	err := fmt.Errorf(format, v...)
	log.Println("❌", err)
	return err
}
