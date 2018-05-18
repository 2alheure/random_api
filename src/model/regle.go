package model

type Regle struct {
	Id         int64       `json:"id"`
	Nom        string      `json:"nom"`
	Parametres []Parametre `json:"parametres"`
}
