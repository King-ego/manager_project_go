package dto

type CreateClassesDTO struct {
	ClassName    string `json:"class_name"`
	TeacherID    string `json:"teacher_id"`
	AcademicYear int    `json:"academic_year"`
	Semester     int    `json:"semester"`
	RoomNumber   int    `json:"room_number"`
}
