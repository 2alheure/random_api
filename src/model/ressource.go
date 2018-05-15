package model

type Ressource struct {
	Id           int64
	Nom          string
	Createur     string
	DateCreation string
	champs       []Champ
}
