package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	ID             string `gorm:"primaryKey; unique; type: uuid; default: uuid_generate_v4()" json:"id"`
	EnrollmentID   string `gorm:"not null;type:uuid" json:"enrollment_id"`
	ScoreValue     int    `gorm:"not null;type:int" json:"score_value"`
	AssessmentType string `gorm:"not null;type:text" json:"assessment_type"`
	AssessmentDate string `gorm:"not null;type:text" json:"assessment_date"`
	Weight         int    `gorm:"not null;type:int" json:"weight"`
	Comments       string `gorm:"type:text" json:"comments"`
}
