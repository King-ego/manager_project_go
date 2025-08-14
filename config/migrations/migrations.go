package migrations

import (
	"log"

	"gorm.io/gorm"
)

type Migration interface {
	Up(*gorm.DB) error
	Down(*gorm.DB) error
}

var allMigrations = []Migration{
	&CreateTableStudents{},
}

func RunMigrations(db *gorm.DB) error {
	for _, migration := range allMigrations {
		if err := migration.Up(db); err != nil {
			return err
		}
		log.Println("Migration applied successfully")
	}

	return nil
}

func RollBackMigrations(db *gorm.DB) error {
	for i := len(allMigrations) - 1; i >= 0; i-- {
		if err := allMigrations[i].Down(db); err != nil {
			return err
		}
		log.Println("Migration applied successfully")
	}
	return nil
}
