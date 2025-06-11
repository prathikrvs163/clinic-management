package controllers

import (
	"clinic-app/config"
	"clinic-app/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	var p models.Patient
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	err = config.DB.QueryRow("INSERT INTO patients(name, age, disease) VALUES($1, $2, $3) RETURNING id",
		p.Name, p.Age, p.Disease).Scan(&p.ID)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func GetPatients(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, age, disease FROM patients")
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var p models.Patient
		err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Disease)
		if err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		patients = append(patients, p)
	}

	json.NewEncoder(w).Encode(patients)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Patient
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	_, err = config.DB.Exec("UPDATE patients SET name=$1, age=$2, disease=$3 WHERE id=$4",
		p.Name, p.Age, p.Disease, id)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM patients WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Delete failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
