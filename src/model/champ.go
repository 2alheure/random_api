package model

import (
	// "encoding/json"
	help "random_api/src/helper"
)

type Champ struct {
	Id    int    `json:"id,omitempty"`
	Clef  string `json:"clef,omitempty"`
	Regle *Regle  `json:"regle,omitempty"`
}

func (champ *Champ) Create() {
	stmt, err := Bdd.Prepare("INSERT INTO champ (clef) VALUE (?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(champ.Clef)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	champ.Id = int(id)
}

func (champ *Champ) Delete() bool {
	stmt, err := Bdd.Prepare("DELETE FROM champ WHERE id=?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(champ.Id)
	help.CheckErr(err)

	affect, err := reponse.RowsAffected()
	help.CheckErr(err)

	if affect != 1 {
		return false
	}
	return true
}

// func (champ *Champ) Generate() interface{} {
// // On va créer un champ en appelant regle.Generate
// // return struct{Key string, Value @type} {"clef", @value}
// 	return struct{}
// }

func GetChamps(max int) []Champ {
	var champs []Champ

	stmt, err := Bdd.Query("SELECT id, clef FROM champ LIMIT ?", max)
	help.CheckErr(err)

	for stmt.Next() {
		var id int
		var clef string
		err = stmt.Scan(&id, &clef)
		help.CheckErr(err)

		champs = append(champs, Champ{Id: id, Clef: clef})
	}

	return champs
}

func GetChamp(id_look int) Champ {
	req := "SELECT id, clef FROM champ WHERE champ.id = ?"
	reponse, err := Bdd.Query(req, id_look)
	defer reponse.Close()
	help.CheckErr(err)

	reponse.Next()

	var id int
	var clef string
	err = reponse.Scan(&id, &clef)
	help.CheckErr(err)

	champ := Champ{Id: id, Clef: clef}
	
	// champ.Hydrate()
	return champ
}

// func (champ *Champ) Hydrate() {
// 	req := "SELECT champ.clef AS clef, champ.id AS champ_id, regle.nom AS regle, regle.id AS regle_id, CONCAT('[', GROUP_CONCAT(CONCAT('{\"id\": ', champ_parametre.id, ', \"type\": \"', parametre.nom, '\", \"value\": \"', champ_parametre.valeur, '\"}') ORDER BY regle_parametre.id), ']') AS parametres FROM champ LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id WHERE champ.id = ? GROUP BY champ_parametre.champ_id"
	
// 	stmt, err := Bdd.Query(req, champ.Id)
// 	defer stmt.Close()
// 	help.CheckErr(err)

	
// 	for stmt.Next() {
// 		var champ_id, regle_id int64
// 		var clef, regle, parametres string
// 		var rule Regle
// 		var params []Parametre
		
// 		err = stmt.Scan(&clef, &champ_id, &regle, &regle_id, &parametres)
// 		help.CheckErr(err)
		
// 		err = json.Unmarshal([]byte(parametres), &params)
// 		help.CheckErr(err)

// 		rule = Regle{
// 			Id: int(regle_id),
// 			Nom: regle,
// 			Parametres: params,
// 		}

// 		champ.Champs = append(champ.Champs, Champ{
// 			Id: int(champ_id),
// 			Clef: clef,
// 			Regle: &rule,
// 		})
// 	}
// }

