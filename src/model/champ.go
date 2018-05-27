package model

import (
	// "encoding/json"
	"net/url"
	help "random_api/src/helper"
)

type Champ struct {
	Id    			int    	`json:"id,omitempty"`
	Clef  			string 	`json:"clef,omitempty"`
	RessourceId    	int		`json:"ressource_id,omitempty"`
	Regle 			*Regle	`json:"regle,omitempty"`
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

func (champ *Champ) Modify(form url.Values) {
	for key, val := range form {
		switch key {
			case "id_ressource":
				stmt, err := Bdd.Prepare("UPDATE champ SET id_ressource = ? WHERE id = ?")
				help.CheckErr(err)
				
				_, err = stmt.Exec(val, champ.Id)
				help.CheckErr(err)
				break;
			case "clef":
				stmt, err := Bdd.Prepare("UPDATE champ SET clef = ? WHERE id = ?")
				help.CheckErr(err)
				
				_, err = stmt.Exec(val, champ.Id)
				help.CheckErr(err)
				break;
		}
	}
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
	
	return champ
}

func (champ *Champ) Hydrate() {
	req := "SELECT DISTINCT champ.ressource_id AS ressource_id, regle_parametre.regle_id AS regle_id FROM champ INNER JOIN champ_parametre ON champ_parametre.champ_id = champ.id INNER JOIN regle_parametre ON regle_parametre.id = champ_parametre.regle_parametre_id WHERE champ.id = ?"
	
	stmt, err := Bdd.Query(req, champ.Id)
	defer stmt.Close()
	help.CheckErr(err)

	
	stmt.Next()
	var ressource_id, regle_id int64
	
	err = stmt.Scan(&ressource_id, &regle_id)
	help.CheckErr(err)
	
	champ.RessourceId = int(ressource_id)
	regle := GetRegle(int(regle_id))
	champ.Regle = &regle
}


