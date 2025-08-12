package config

import (
	"database/sql"
	"fmt"
)

const (
	host     = "manager_db"
	port     = 5432
	user     = "manager"
	password = "manager"
	dbname   = "manager_db"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to open database connection: " + err.Error())
	}

	fmt.Println("connected to database " + dbname)

	return db, nil
}
