package main

import (
	"log"
	"manager_project/config"
	"manager_project/config/migrations"
)

func main() {
	db, err := config.ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Starting rollback...")
	if err := migrations.RollBackMigrations(db); err != nil {
		log.Fatalf("Could not rollback migrations: %v", err)
	}
	log.Println("Rollback completed successfully")
}
