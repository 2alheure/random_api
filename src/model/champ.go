package model

import help "random_api/src/helper"

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

	stmt, err := Bdd.Query("SELECT * FROM champ LIMIT ?", max)
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
