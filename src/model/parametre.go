package model

type Parametre struct {
	Id    int    `json:"id,omitempty"`
	Nom   string `json:"nom,omitempty"`
	Value string `json:"value,omitempty"`
}
