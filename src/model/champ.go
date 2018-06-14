package model

import (
	// "encoding/json"
	"net/url"
	"gopkg.in/guregu/null.v3"
	
	_ "fmt"
	help "../helper"
)

type Champ struct {
	Id    			int    			`json:"id,omitempty"`
	Clef  			null.String 	`json:"clef,omitempty"`
	RessourceId    	null.Int		`json:"ressource_id,omitempty"`
	Regle 			*Regle			`json:"regle,omitempty"`
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

func (champ *Champ) Modify(form url.Values) bool {
	for key, val := range form {
		switch key {
			case "ressource_id":
				var ress_id string
				if val[0] == "" {
					ress_id = "NULL"
				} else {
					ress_id = val[0]
				}

				stmt, err := Bdd.Prepare("UPDATE champ SET ressource_id = ? WHERE id = ?")
				help.CheckErr(err)
				
				reponse, err := stmt.Exec(ress_id, champ.Id)
				help.CheckErr(err)

				affect, err := reponse.RowsAffected()
				help.CheckErr(err)
			
				if affect != 1 {
					return false
				}
				break;
			case "clef":
				stmt, err := Bdd.Prepare("UPDATE champ SET clef = ? WHERE id = ?")
				help.CheckErr(err)
				
				reponse, err := stmt.Exec(val[0], champ.Id)
				help.CheckErr(err)
				
				affect, err := reponse.RowsAffected()
				help.CheckErr(err)
			
				if affect != 1 {
					return false
				}
				break;
		}
	}
	return true;
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

func GetChamps(max int) []Champ {
	var champs []Champ

	stmt, err := Bdd.Query("SELECT id, clef, ressource_id FROM champ LIMIT ?", max)
	help.CheckErr(err)

	for stmt.Next() {
		var champ Champ
		err = stmt.Scan(&champ.Id, &champ.Clef, &champ.RessourceId)
		help.CheckErr(err)
	
		champs = append(champs, champ)
	}

	return champs
}

func GetChamp(id_look int) Champ {
	req := "SELECT id, clef FROM champ WHERE champ.id = ?"
	reponse, err := Bdd.Query(req, id_look)
	defer reponse.Close()
	help.CheckErr(err)

	reponse.Next()

	var champ Champ
	err = reponse.Scan(&champ.Id, &champ.Clef)
	help.CheckErr(err)

	return champ
}

func (champ *Champ) Hydrate() {
	req := "SELECT DISTINCT champ.ressource_id AS ressource_id, regle_parametre.regle_id AS regle_id FROM champ INNER JOIN champ_parametre ON champ_parametre.champ_id = champ.id INNER JOIN regle_parametre ON regle_parametre.id = champ_parametre.regle_parametre_id WHERE champ.id = ?"
	
	stmt, err := Bdd.Query(req, champ.Id)
	defer stmt.Close()
	help.CheckErr(err)

	
	stmt.Next()
	var ressource_id, regle_id null.Int
	
	err = stmt.Scan(&ressource_id, &regle_id)
	help.CheckErr(err)
	
	champ.RessourceId = ressource_id

	if regle_id.Valid {
		regle := GetRegle(int(regle_id.Int64))
		champ.Regle = &regle
	}
}
