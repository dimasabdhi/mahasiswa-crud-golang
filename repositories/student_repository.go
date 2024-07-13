package repositories

import (
	"student-crud/config"
	"student-crud/models"
)

func GetAllStudents() []models.Student {
	var students []models.Student
	config.DB.Find(&students)
	return students
}

func GetStudentByNPM(npm string) (models.Student, error) {
	var student models.Student
	result := config.DB.First(&student, "npm = ?", npm)
	return student, result.Error
}

func CreateStudent(student models.Student) models.Student {
	config.DB.Create(&student)
	return student
}

func UpdateStudent(student models.Student) models.Student {
	config.DB.Save(&student)
	return student
}

func DeleteStudent(npm string) error {
	result := config.DB.Delete(&models.Student{}, "npm = ?", npm)
	return result.Error
}
