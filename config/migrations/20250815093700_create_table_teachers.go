package migrations

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type CreateTableTeachers struct{}

func (m *CreateTableTeachers) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Teacher{})
}

func (m *CreateTableTeachers) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Teacher{})
}
