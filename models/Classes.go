package models

import (
	"time"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	ID        string    `gorm:"primaryKey;default:uuid_generate_v4();type:uuid" json:"id"`
	TeacherID string    `gorm:"type:uuid" json:"teacher_id"`
	Year      int       `gorm:"type:int" json:"year"`
	Semester  int       `gorm:"type:int" json:"semester"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
