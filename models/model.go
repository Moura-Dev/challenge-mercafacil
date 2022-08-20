package models

type ContactInfo struct {
	ID      string `db:"id" json:"id"`
	Nome    string `db:"nome" json:"name"`
	Celular string `db:"celular" json:"cellphone"`
}

type Contacts struct {
	Infos []ContactInfo `json:"contacts"`
}
