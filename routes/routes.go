package routes

import (
	"clinic-app/controllers"
	"clinic-app/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	// Auth
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Patient Routes
	r.HandleFunc("/patients", middleware.AuthMiddleware(controllers.CreatePatient, "receptionist")).Methods("POST")
	r.HandleFunc("/patients", middleware.AuthMiddleware(controllers.GetPatients, "receptionist", "doctor")).Methods("GET")
	r.HandleFunc("/patients/{id}", middleware.AuthMiddleware(controllers.UpdatePatient, "doctor")).Methods("PUT")
	r.HandleFunc("/patients/{id}", middleware.AuthMiddleware(controllers.DeletePatient, "receptionist")).Methods("DELETE")

	return r
}
