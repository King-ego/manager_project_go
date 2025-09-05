package migrations

import (
	"log"
	"reflect"
	"time"

	"gorm.io/gorm"
)

type Migration interface {
	Up(*gorm.DB) error
	Down(*gorm.DB) error
}

type MigrationRecord struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	ExecutedAt time.Time
}

var allMigrations = []Migration{
	&CreateTableStudents{},
	&CreateTableClasses{},
	&CreateTableScores{},
	&CreateTableEnrollments{},
	&CreateTableTeachers{},
}

func RunMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&MigrationRecord{}); err != nil {
		log.Printf("Failed to create migration_records table: %v", err)
		return err
	}

	for _, migration := range allMigrations {
		migrationName := reflect.TypeOf(migration).Elem().Name()

		var record MigrationRecord
		if err := db.Where("name = ?", migrationName).First(&record).Error; err == nil {
			log.Printf("Migration %s already executed, skipping...", migrationName)
			continue
		}

		log.Printf("Running migration: %s", migrationName)
		if err := migration.Up(db); err != nil {
			return err
		}

		if err := db.Create(&MigrationRecord{
			Name:       migrationName,
			ExecutedAt: time.Now(),
		}).Error; err != nil {
			return err
		}

		log.Printf("Migration %s applied successfully", migrationName)
	}

	return nil
}

func RollBackMigrations(db *gorm.DB) error {
	if !db.Migrator().HasTable(&MigrationRecord{}) {
		log.Println("Migration records table doesn't exist, nothing to rollback")
		return nil
	}

	var lastRecord MigrationRecord
	if err := db.Order("executed_at DESC").First(&lastRecord).Error; err != nil {
		log.Println("No migrations found to rollback")
		return nil
	}

	var targetMigration Migration
	for _, migration := range allMigrations {
		migrationName := reflect.TypeOf(migration).Elem().Name()
		if migrationName == lastRecord.Name {
			targetMigration = migration
			break
		}
	}

	if targetMigration == nil {
		log.Printf("Migration %s not found in code, cannot rollback", lastRecord.Name)
		return nil
	}

	log.Printf("Rolling back last migration: %s", lastRecord.Name)
	if err := targetMigration.Down(db); err != nil {
		return err
	}

	if err := db.Where("name = ?", lastRecord.Name).Delete(&MigrationRecord{}).Error; err != nil {
		return err
	}

	log.Printf("Migration %s rolled back successfully", lastRecord.Name)
	return nil
}
