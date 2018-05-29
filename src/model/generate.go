package model

import (
	_ "fmt"
)

type RessourceKV struct {
	Champs		[]ChampKV		`json:",omitempty"`
}

type ChampKV struct {
	Int		int
	Float	float64
	String	string
	Bool	bool
}


var RuleSet = map[int]					// Doit servir à savoir combien de paramètres fournir aux règles (need une fonction pour récup toutes les règles de la base ?)

/* Useful :
	https://gobyexample.com/variadic-functions

*/


func Generate(ressource_id, nombre int) (map[string]ChampKV, int) {
	ressource, err := GetRessource(ressource_id)
	if err != nil {
		return nil, 404
	}

	ressource.Hydrate()

	ret := make(map[string]ChampKV)

	for _, champ := range ressource.Champs {

		if !champ.Clef.Valid || champ.Regle == nil  || champ.Regle.Id == 0  {
			return ret, 409
		}

		params := champ.Regle.Parametres

		if len(params) == 0 {
			return ret, 409
		} else {

			for _, param := range param {

				if param.Type == "" || param.Value == "" {
					return ret, 409
				} else {

				}
			}
		}
	}















	return ret, 0
}