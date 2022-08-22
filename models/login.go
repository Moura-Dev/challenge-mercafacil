package models

type Login struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Customer string `db:"customer" json:"customer"`
}

// Create Check User func
