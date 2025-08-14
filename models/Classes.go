package models

import (
	"time"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	ID           string    `gorm:"primaryKey;unique;default:uuid_generate_v4();type:uuid" json:"id"`
	ClassName    string    `gorm:"not null" json:"className"`
	TeacherID    string    `gorm:"type:uuid" json:"teacher_id"`
	AcademicYear int       `gorm:"type:int" json:"academic_year"`
	Semester     int       `gorm:"type:int" json:"semester"`
	RoomNumber   int       `gorm:"type:int" json:"room_number"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
