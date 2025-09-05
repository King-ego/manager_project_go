package migrations

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type CreateTableScores struct{}

func (m *CreateTableScores) Up(db *gorm.DB) error {
	return db.AutoMigrate(&models.Score{})
}

func (m *CreateTableScores) Down(db *gorm.DB) error {
	return db.Migrator().DropTable(&models.Score{})
}
