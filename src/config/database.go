package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connects to the PostgreSQL database using GORM.
// Returns a pointer to the GORM DB object and any error encountered.
func ConnectDB() (*gorm.DB, error) {
	dbURL := getDBURL("DATABASE_URL")

	// Attempt to open a connection to the database.
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		// Handle error (for example, log it and return it)
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	fmt.Println("Connected to the database successfully.")
	return db, nil
}

func getDBURL(urlEnv string) string {
	dbURL := os.Getenv(urlEnv)
	if dbURL == "" {
		// Use the service name as the hostname
		dbURL = "host=lsplanner-database user=lsplanner password=lsplanner dbname=lsplanner port=5432 sslmode=disable"
	}
	return dbURL
}
