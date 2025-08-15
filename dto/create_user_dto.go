package dto

type CreateUserDTO struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required,datetime=2006-01-02"`
}
