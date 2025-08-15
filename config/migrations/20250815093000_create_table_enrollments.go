package migrations

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type CreateTableEnrollments struct {
}

func (m *CreateTableEnrollments) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Enrollment{})
}

func (m *CreateTableEnrollments) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Enrollment{})
}
