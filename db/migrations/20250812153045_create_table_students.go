package migrations

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type CreateTableStudents struct{}

func (m *CreateTableStudents) Up(db *gorm.DB) error {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	return db.AutoMigrate(&models.Students{})
}

func (m *CreateTableStudents) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Students{})
}
