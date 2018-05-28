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

func GetRegle(id int) Regle {
	stmt, err := Bdd.Query("SELECT regle.nom AS regle, regle.id AS regle_id, CONCAT( '[', GROUP_CONCAT( CONCAT( '{\"id\": ', parametre.id, ', ' '\"type\": \"', parametre.nom, '\"}' ) ORDER BY regle_parametre.id ), ']' ) AS parametres FROM regle LEFT OUTER JOIN regle_parametre ON regle.id = regle_parametre.regle_id LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id WHERE regle.id = ? GROUP BY regle_parametre.regle_id", id)
	help.CheckErr(err)

	stmt.Next()

	var regle_id int64
	var nom, params string
	var parametres []Parametre
	err = stmt.Scan(&nom, &regle_id, &params)
	help.CheckErr(err)

	err = json.Unmarshal([]byte(params), &parametres)
	help.CheckErr(err)

	regle := Regle{
		Id: int(regle_id),
		Nom: nom,
		Parametres: parametres,
	}

	return regle
}

func AttachRule(regle_id, champ_id int) bool {
	stmt, err := Bdd.Prepare("INSERT INTO champ_parametre (champ_id, regle_parametre_id) SELECT champ.id AS champ, t.id FROM champ CROSS JOIN (SELECT id FROM regle_parametre WHERE regle_parametre.regle_id = ?) AS t WHERE champ.id = ?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(regle_id, champ_id)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	if id < 1 {
		return false
	}
	return true
}

func DetachRule(champ_id int) bool {
	stmt, err := Bdd.Prepare("DELETE FROM champ_parametre WHERE champ_id=?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(champ_id)
	help.CheckErr(err)

	affect, err := reponse.RowsAffected()
	help.CheckErr(err)

	if affect < 1 {
		return false
	}
	return true
}
