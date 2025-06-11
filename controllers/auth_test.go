package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"clinic-app/config"
)

// Setup a known user in DB before running tests
func setupTestUser() {
	config.Connect() // make sure DB is connected

	// Delete existing user if exists
	config.DB.Exec("DELETE FROM users WHERE email = 'testuser@example.com'")

	// Insert test user with bcrypt hashed password
	hashed := "$2a$10$ff35NXJrXmFZFlaOm7XcfOaiPiFZPpEksGkGC9xEyGdG8fzwvCNOq"
	config.DB.Exec(`INSERT INTO users (email, password, role) VALUES ($1, $2, $3)`,
		"testuser@example.com", hashed, "receptionist")
}

func TestLoginSuccess(t *testing.T) {
	setupTestUser()

	payload := map[string]string{
		"email":    "testuser@example.com",
		"password": "test123",
	}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", rr.Code)
	}

	// Optionally check that response contains a token
	var res map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	if err != nil {
		t.Fatal("Invalid JSON response")
	}
	if res["token"] == "" {
		t.Error("Expected a token in response")
	}
}
