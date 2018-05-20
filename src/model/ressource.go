package model

import (
	help "../helper"
)

type Ressource struct {
	Id           int     `json:"id"`
	Nom          string  `json:"nom"`
	Createur     string  `json:"createur"`
	DateCreation string  `json:"date_creation"`
	Champs       []Champ `json:"champs"`
}

func (ress *Ressource) Create() {
	stmt, err := db.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(ress.Nom, ress.Createur)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	ret := GetRessource(int(id))
	ress = &ret
}

func (ress *Ressource) Delete() bool {
	stmt, err := db.Prepare("DELETE FROM ressource WHERE id=?")
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
// 	var ret = []Ressource

// 	for i := 0 ; i<n ; i++ {
// 		for _, j := range ress.Champs {
// 			ret = append(ret, j.Generate())
// 		}
// 	}

// 	return ret
// }

func GetRessources(max int) []Ressource {
	stmt, err := db.Query("SELECT * FROM ressource LIMIT ?", max)
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
	reponse, err := db.Query("SELECT ressource.*, champ.clef FROM ressource LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id WHERE ressource.id = ? ORDER BY clef", id_look)
	help.CheckErr(err)
	defer reponse.Close()

	reponse.Next()

	var id int64
	var nom, createur, date, clef string
	err = reponse.Scan(&id, &nom, &createur, &date, &clef)
	help.CheckErr(err)
	
	ress := Ressource{
		Id: int(id), 
		Nom: nom, 
		Createur: createur, 
		DateCreation: date,
		Champs: make([]Champ, 0),
	}

	ress.Champs = append(ress.Champs, Champ{Clef: clef})

	for reponse.Next() {
		err = reponse.Scan(&id, &nom, &createur, &date, &clef)
		help.CheckErr(err)
		ress.Champs = append(ress.Champs, Champ{Clef: clef})
	}

	return ress
}

// func (ress *Ressource) Hydrate() {
// 	req := "SELECT ressource.*, champ.clef FROM ressource LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id WHERE ressource.id = ? ORDER BY clef"
	
// 	stmt, err := db.Query(req, id)
// 	help.CheckErr(err)

// 	stmt.Next()
// 	ret := Ressource{Id: }
// }

