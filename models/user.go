package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"` // hashed
	Role     string `json:"role"`     // doctor or receptionist
}
