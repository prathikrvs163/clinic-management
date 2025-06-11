package main

import (
	"clinic-app/config"
	"clinic-app/middleware"
	"clinic-app/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Connect()

	// Optional: Create patients table
	_, err := config.DB.Exec(`
		CREATE TABLE IF NOT EXISTS patients (
			id SERIAL PRIMARY KEY,
			name TEXT,
			age INT,
			disease TEXT
		);
	`)
	if err != nil {
		log.Fatal("Error creating patients table:", err)
	}

	// Optional: Create users table
	_, err = config.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			email TEXT PRIMARY KEY,
			password TEXT NOT NULL,
			role TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	r := routes.RegisterRoutes()

	// Wrap routes with CORS middleware
	handler := middleware.CORSMiddleware(r)

	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
