package models

type Employee struct {
	ID         int        `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	Department Department `json:"department"`
}
