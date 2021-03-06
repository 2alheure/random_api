package model

import (
	"encoding/json"
	"net/url"
	"gopkg.in/guregu/null.v3"

	help "../helper"
)

type Ressource struct {
	Id           int     		`json:"id,omitempty"`
	Nom          null.String  	`json:"nom,omitempty"`
	Createur     null.String  	`json:"createur,omitempty"`
	DateCreation null.String  	`json:"date_creation,omitempty"`
	Champs       []*Champ 		`json:"champs,omitempty"`
}

func (ress *Ressource) Create() {
	stmt, err := Bdd.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(ress.Nom, ress.Createur)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	ressource, err := GetRessource(int(id))
	help.CheckErr(err)

	ress.DateCreation = ressource.DateCreation
	ress.Id = ressource.Id
}

func (ress *Ressource) Modify(form url.Values) {
	for key, val := range form {
		switch key {
			case "nom":
				stmt, err := Bdd.Prepare("UPDATE ressource SET nom = ? WHERE id = ?")
				help.CheckErr(err)
				
				_, err = stmt.Exec(val[0], ress.Id)
				help.CheckErr(err)
				break;
			case "createur":
				stmt, err := Bdd.Prepare("UPDATE ressource SET createur = ? WHERE id = ?")
				help.CheckErr(err)
				
				_, err = stmt.Exec(val[0], ress.Id)
				help.CheckErr(err)
				break;
		}
	}
}

func (ress *Ressource) Delete() bool {
	stmt, err := Bdd.Prepare("DELETE FROM ressource WHERE id=?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(ress.Id)
	help.CheckErr(err)

	affect, err := reponse.RowsAffected()
	help.CheckErr(err)

	if affect != 1 {
		return false
	}
	return true
}

// func (ress *Ressource) Generate(n int) []Ressource {
// // On va créer un array qui récupère les clefs de chaque champ et leur associe une valeur
// // La valeur de chaque champ vaudra champ.Generate
// 	var ret = []map[string]Champ

// 	for i := 0 ; i<n ; i++ {
// 		for _, j := range ress.Champs {
// 			ret = append(ret, j.Generate())
// 		}
// 	}

// 	return ret
// }

func GetRessources(max int) []Ressource {
	stmt, err := Bdd.Query("SELECT * FROM ressource ORDER BY date_creation DESC LIMIT ?", max)
	help.CheckErr(err)
	
	var ress []Ressource
	
	for stmt.Next() {
		var ressource Ressource
		err = stmt.Scan(&ressource.Id, &ressource.Nom, &ressource.Createur, &ressource.DateCreation)
		help.CheckErr(err)

		ress = append(ress, ressource)
	}

	return ress
}

func GetRessource(id_look int) (Ressource, error) {
	req := "SELECT * FROM ressource WHERE ressource.id = ?"
	reponse, err := Bdd.Query(req, id_look)
	defer reponse.Close()
	help.CheckErr(err)

	reponse.Next()

	var ress Ressource
	err = reponse.Scan(&ress.Id, &ress.Nom, &ress.Createur, &ress.DateCreation)
	help.CheckErr(err)
	
	return ress, err
}

func (ress *Ressource) Hydrate() {
	req := "SELECT champ.clef AS clef, champ.id AS champ_id, regle.nom AS regle, regle.id AS regle_id, CONCAT('[', GROUP_CONCAT(CONCAT('{\"id\": ', champ_parametre.id, ', \"type\": \"', parametre.nom, '\", \"value\": \"', champ_parametre.valeur, '\"}') ORDER BY regle_parametre.id), ']') AS parametres FROM champ LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id WHERE champ.ressource_id = ? GROUP BY champ.id ORDER BY clef"
	
	stmt, err := Bdd.Query(req, ress.Id)
	defer stmt.Close()
	help.CheckErr(err)

	
	for stmt.Next() {
		var champ_id, regle_id null.Int
		var clef, regle, parametres null.String
		var rule Regle
		var params []Parametre
		
		err = stmt.Scan(&clef, &champ_id, &regle, &regle_id, &parametres)
		help.CheckErr(err)
		
		err = json.Unmarshal([]byte(parametres.String), &params)
		help.CheckErr(err)

		rule = Regle{
			Id: int(regle_id.Int64),
			Nom: regle.String,
			Parametres: params,
		}

		ress.Champs = append(ress.Champs, &Champ{
			Id: int(champ_id.Int64),
			Clef: clef,
			Regle: &rule,
		})
	}
}

