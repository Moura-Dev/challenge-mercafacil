package models

type User struct {
	ID       string `db:"id" json:"id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Customer string `db:"customer" json:"customer"`
}

// Create Check User func
