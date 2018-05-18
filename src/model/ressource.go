package model

type Ressource struct {
	Id           int64   `json:"id"`
	Nom          string  `json:"nom"`
	Createur     string  `json:"createur"`
	DateCreation string  `json:"date_creation"`
	Champs       []Champ `json:"champs"`
}
