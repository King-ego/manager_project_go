package models

import "gorm.io/gorm"

type Enrollment struct {
	gorm.Model
	ID        string `gorm:"primaryKey; unique; type:uuid; default:uuid_generate_v4()" json:"id"`
	StudentID string `gorm:"not null; type:uuid;" json:"student_id"`
	ClassID   string `gorm:"not null; type:uuid;" json:"class_id"`
	Status    bool   `gorm:"not null; type:boolean; default: true" json:"status"`
}
