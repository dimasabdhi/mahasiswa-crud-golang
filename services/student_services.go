package services

import (
	"student-crud/models"
	"student-crud/repositories"
)

func GetAllStudents() []models.Student {
	return repositories.GetAllStudents()
}

func GetStudentByNPM(npm string) (models.Student, error) {
	return repositories.GetStudentByNPM(npm)
}

func CreateStudent(student models.Student) models.Student {
	return repositories.CreateStudent(student)
}

func UpdateStudent(student models.Student) models.Student {
	return repositories.UpdateStudent(student)
}

func DeleteStudent(npm string) error {
	return repositories.DeleteStudent(npm)
}
