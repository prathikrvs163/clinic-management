package models

type Patient struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Disease string `json:"disease"`
}
