package model

import (
	help "random_api/src/helper"
	"encoding/json"
)

type Regle struct {
	Id         int         `json:"id,omitempty"`
	Nom        string      `json:"nom,omitempty"`
	Parametres []Parametre `json:"parametres,omitempty"`
}


func (regle *Regle) Create() {
	stmt, err := Bdd.Prepare("INSERT INTO regle (Nom) VALUE (?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(regle.Nom)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	regle.Id = int(id)
}

func (regle *Regle) Delete() bool {
	stmt, err := Bdd.Prepare("DELETE FROM regle WHERE id=?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(regle.Id)
	help.CheckErr(err)

	affect, err := reponse.RowsAffected()
	help.CheckErr(err)

	if affect != 1 {
		return false
	}
	return true
}

func GetRegles(max int) []Regle {
	var regles []Regle

	stmt, err := Bdd.Query("SELECT regle.nom AS regle, regle.id AS regle_id, CONCAT( '[', GROUP_CONCAT( CONCAT( '{\"id\": ', parametre.id, ', ' '\"type\": \"', parametre.nom, '\"}' ) ORDER BY regle_parametre.id ), ']' ) AS parametres FROM regle LEFT OUTER JOIN regle_parametre ON regle.id = regle_parametre.regle_id LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id GROUP BY regle_parametre.regle_id LIMIT ?", max)
	help.CheckErr(err)

	for stmt.Next() {
		var id int64
		var nom, params string
		var parametres []Parametre
		err = stmt.Scan(&nom, &id, &params)
		help.CheckErr(err)

		err = json.Unmarshal([]byte(params), &parametres)
		help.CheckErr(err)

		rule := Regle{
			Id: int(id),
			Nom: nom,
			Parametres: parametres,
		}

		regles = append(regles, rule)
	}

	return regles
}
