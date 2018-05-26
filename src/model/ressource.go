package model

import (
	"encoding/json"
	help "random_api/src/helper"
)

type Ressource struct {
	Id           int     `json:"id,omitempty"`
	Nom          string  `json:"nom,omitempty"`
	Createur     string  `json:"createur,omitempty"`
	DateCreation string  `json:"date_creation,omitempty"`
	Champs       []Champ `json:"champs,omitempty"`
}

func (ress *Ressource) Create() {
	stmt, err := Bdd.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(ress.Nom, ress.Createur)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	ret := GetRessource(int(id))
	ress = &ret
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
	stmt, err := Bdd.Query("SELECT * FROM ressource LIMIT ?", max)
	help.CheckErr(err)
	
	var ress []Ressource
	
	for stmt.Next() {
		var id int64
		var nom, createur, date string
		err = stmt.Scan(&id, &nom, &createur, &date)
		help.CheckErr(err)

		ress = append(ress, Ressource{
			Id: int(id), 
			Nom: nom, 
			Createur: createur, 
			DateCreation: date,
		})
	}

	return ress
}

func GetRessource(id_look int) Ressource {
	req := "SELECT * FROM ressource WHERE ressource.id = ?"
	reponse, err := Bdd.Query(req, id_look)
	defer reponse.Close()
	help.CheckErr(err)

	reponse.Next()

	var id int64
	var nom, createur, date string
	err = reponse.Scan(&id, &nom, &createur, &date)
	help.CheckErr(err)
	
	ress := Ressource{
		Id: int(id), 
		Nom: nom, 
		Createur: createur, 
		DateCreation: date,
	}

	ress.Hydrate()
	return ress
}

func (ress *Ressource) Hydrate() {
	req := "SELECT champ.clef AS clef, champ.id AS champ_id, regle.nom AS regle, regle.id AS regle_id, GROUP_CONCAT(CONCAT('{\"id\": \"', champ_parametre.id, '\", \"type\": \"', parametre.nom, '\", \"value\": \"', champ_parametre.valeur, '\"}') ORDER BY regle_parametre.id) AS parametres FROM ressource LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id WHERE ressource.id = ? GROUP BY champ_parametre.champ_id ORDER BY clef"
	
	stmt, err := Bdd.Query(req, ress.Id)
	defer stmt.Close()
	help.CheckErr(err)

	var champ_id, regle_id int64
	var clef, regle, parametres string
	var rule Regle
	var params []Parametre

	for stmt.Next() {
		err = stmt.Scan(&clef, &champ_id, &regle, &regle_id, &parametres)
		help.CheckErr(err)
		
		err = json.Unmarshal([]byte(parametres), &params)
		help.CheckErr(err)

		rule = Regle{
			Id: int(regle_id),
			Nom: regle,
			Parametres: params,
		}

		ress.Champs = append(ress.Champs, Champ{
			Id: int(champ_id),
			Clef: clef,
			Regle: &rule,
		})
	}
}

