package main

import (
	"log"
	"net/http"
	"student-crud/config"
	"student-crud/routes"

	"github.com/rs/cors"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Create a new router
	router := routes.SetupRouter()

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Use the CORS middleware
	handler := c.Handler(router)

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
