package model

import (
	help "../helper"
)

type Regle struct {
	Id         int         `json:"id"`
	Nom        string      `json:"nom"`
	Parametres []Parametre `json:"parametres"`
}


func (regle *Regle) Create() {
	stmt, err := db.Prepare("INSERT INTO regle (Nom) VALUE (?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(regle.Nom)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	regle.Id = int(id)
}

func (regle *Regle) Delete() bool {
	stmt, err := db.Prepare("DELETE FROM regle WHERE id=?")
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

	stmt, err := db.Query("SELECT * FROM regle LIMIT ?")
	help.CheckErr(err)

	for stmt.Next() {
		var id int
		var nom string
		err = stmt.Scan(&id, &nom)
		help.CheckErr(err)

		regles = append(regles, Regle{Id: id, Nom: nom})
	}

	return regles
}
