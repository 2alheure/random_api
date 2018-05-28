package model

type Parametre struct {
	Id    int    `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func SetParametres(champ_id int, parametres []string) {
	for _, val := range parametres {
		stmt, err := Bdd.Prepare("INSERT INTO")
		/*
https://stackoverflow.com/questions/25674737/mysql-update-multiple-rows-with-different-values-in-one-query
https://stackoverflow.com/questions/3870540/how-to-update-column-with-null-value
https://stackoverflow.com/questions/2472229/insert-into-select-from-on-duplicate-key-update
		*/
		help.CheckErr(err)
		
		_, err = stmt.Exec(val[0], ress.Id)
		help.CheckErr(err)
	}
}

func ResetParametres(champ_id int) {
	stmt, err := Bdd.Prepare("UPDATE champ_parametre SET valeur = NULL WHERE champ_id = ?")
	help.CheckErr(err)
	
	ret, err := stmt.Exec(champ_id)
	help.CheckErr(err)

	if ret < 1 {
		return false
	}
	return true
}

