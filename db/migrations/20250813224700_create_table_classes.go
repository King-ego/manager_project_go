package migrations

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type CreateTableClasses struct{}

func (m *CreateTableClasses) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Classes{})
}

func (m *CreateTableClasses) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Classes{})
}
