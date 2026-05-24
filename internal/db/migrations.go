package db

import (
	"database/sql"
	"os"
	"log"
	"fmt"
)

func RunMigrations(db *sql.DB) {
	files := []string{
		"db/migrations/001_create_users.sql",
		"db/migrations/002_create_transactions.sql",
	}

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Error reading migration file: ", err)
		}

		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal("Error running migration: ", err)
		}

		fmt.Println("Ran migration: ", file)
	}
}