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

	log.Println("Starting migrations...")
	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
