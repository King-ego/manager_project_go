package models

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID             string    `gorm:"primaryKey;unique;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name           string    `gorm:"size:100;not null" json:"name"`
	Email          string    `gorm:"size:100;not null" json:"email"`
	Specialization string    `gorm:"size:100;not null" json:"specialization"`
	HireDate       time.Time `gorm:"type:timestamp" json:"hire_date"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
