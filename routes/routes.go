package routes

import (
	"student-crud/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/{npm}", controllers.GetStudentByNPM).Methods("GET")
	router.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/students/{npm}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{npm}", controllers.DeleteStudent).Methods("DELETE")

	return router
}
