package model

type Regle struct {
	Id         int         `json:"id"`
	Nom        string      `json:"nom"`
	Parametres []Parametre `json:"parametres"`
}
