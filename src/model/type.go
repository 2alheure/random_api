package model

type Type struct {
	Id     int64   `json:"id"`
	Nom    string  `json:"nom"`
	Regles []Regle `json:"regles"`
}
