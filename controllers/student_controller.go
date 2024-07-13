package controllers

import (
	"encoding/json"
	"net/http"
	"student-crud/models"
	"student-crud/services"

	"github.com/gorilla/mux"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := services.GetAllStudents()
	json.NewEncoder(w).Encode(students)
}

func GetStudentByNPM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	student, err := services.GetStudentByNPM(params["npm"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)
	student = services.CreateStudent(student)
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var student models.Student
	student.NPM = params["npm"]
	json.NewDecoder(r.Body).Decode(&student)
	student = services.UpdateStudent(student)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := services.DeleteStudent(params["npm"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
