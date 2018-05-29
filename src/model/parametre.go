package model

import (
	help "../helper"
)

type Parametre struct {
	Id    int    `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func SetParametres(champ_id int, parametres []string) bool {
	if lenParam := len(parametres); lenParam > 0 {
		req1 := "SELECT id FROM champ_parametre WHERE champ_id = ?"
		rep1, err := Bdd.Query(req1, champ_id)
		help.CheckErr(err)

		var ids []int64
		var id int64
		for rep1.Next() {
			err = rep1.Scan(&id)
			help.CheckErr(err)

			ids = append(ids, id)
		}

		

		if lenIds := len(ids); lenParam == lenIds {
			req2 := "UPDATE champ_parametre SET valeur = ? WHERE id = ?"
			for i := 0; i < lenParam; i++ {
				stmt, err := Bdd.Prepare(req2)
				help.CheckErr(err)
				
				ret, err := stmt.Exec(parametres[i], ids[i])
				help.CheckErr(err)

				affect, err := ret.RowsAffected()
				help.CheckErr(err)
			
				if affect < 1 {
					return false
				}
			}

			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ResetParametres(champ_id int) bool {
	stmt, err := Bdd.Prepare("UPDATE champ_parametre SET valeur = NULL WHERE champ_id = ?")
	help.CheckErr(err)
	
	ret, err := stmt.Exec(champ_id)
	help.CheckErr(err)

	affect, err := ret.RowsAffected()
	help.CheckErr(err)

	if affect < 1 {
		return false
	}
	return true
}

