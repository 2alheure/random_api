package model

import help "../helper"

type Champ struct {
	Id    int    `json:"id"`
	Clef  string `json:"clef"`
	Regle Regle  `json:"regle"`
}

func (champ *Champ) Create() {
	stmt, err := db.Prepare("INSERT INTO champ (clef) VALUE (?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(champ.Clef)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	champ.Id = int(id)
}

func (champ *Champ) Delete() bool {
	stmt, err := db.Prepare("DELETE FROM champ WHERE id=?")
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

// func (champ *Champ) Generate() Champ {
// On va créer un champ en appelant regle.Generate
// }

func GetChamps(max int) []Champ {
	var champs []Champ

	stmt, err := db.Query("SELECT * FROM champ LIMIT ?")
	help.CheckErr(err)

	for stmt.Next() {
		var id int
		var nom, createur, date string
		err = stmt.Scan(&id, &nom, &createur, &date)
		help.CheckErr(err)

		champs = append(champs, Champ{Id: id, Nom: nom, Createur: createur, DateCreation: date})
	}

	return champs
}
