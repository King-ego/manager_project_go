package main

import (
	"log"
	"manager_project/db"
	"manager_project/db/migrations"
)

func main() {
	database, err := db.ConnectDb()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Starting migrations...")
	if err := migrations.RunMigrations(database); err != nil {
		log.Fatalf("Could not run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}
