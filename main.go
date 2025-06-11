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

	// Create patients table if not exists
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

	// Create users table if not exists
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

	// Register API routes
	r := routes.RegisterRoutes()

	// Serve static frontend from ./public
	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/").Handler(fs) // for index.html, CSS, JS, etc.

	// Wrap everything with CORS middleware
	handler := middleware.CORSMiddleware(r)

	fmt.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
