package dto

import "manager_project/models"

type CreateClassesDTO struct {
	ClassName    string `json:"class_name"`
	TeacherID    string `json:"teacher_id"`
	AcademicYear int    `json:"academic_year"`
	Semester     int    `json:"semester"`
	RoomNumber   int    `json:"room_number"`
}

func (dto *CreateClassesDTO) ToModel() *models.Classes {
	return &models.Classes{
		ClassName:    dto.ClassName,
		TeacherID:    dto.TeacherID,
		AcademicYear: dto.AcademicYear,
		Semester:     dto.Semester,
		RoomNumber:   dto.RoomNumber,
	}
}
