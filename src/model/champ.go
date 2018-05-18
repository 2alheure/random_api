package model

type Champ struct {
	Id   int64  `json:"id"`
	Clef string `json:"clef"`
	Type Type   `json:"type"`
}
