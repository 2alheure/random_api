package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	help "../helper"
)

var db *sql.DB

func TestSql() {
	fmt.Println("test sql")
}

func Connect() {
	var err error
	db, err = sql.Open("mysql", "root:@/alea_data_est?charset=utf8")
	help.CheckErr(err)
}

func CreateRessource(nom, createur string) int64 {
	stmt, err := db.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)")
	help.CheckErr(err)

	reponse, err := stmt.Exec(nom, createur)
	help.CheckErr(err)

	id, err := reponse.LastInsertId()
	help.CheckErr(err)

	return id
}

func GetRessources(max int) []Ressource {
	var res []Ressource

	stmt, err := db.Query("SELECT * FROM ressource LIMIT ?")
	help.CheckErr(err)

	for stmt.Next() {
		var id int64
		var nom, createur, date string
		err = stmt.Scan(&id, &nom, &createur, &date)
		help.CheckErr(err)

		res = append(res, Ressource{Id: id, Nom: nom, Createur: createur, DateCreation: date})
	}

	return res
}

func DeleteRessource(id int) bool {
	stmt, err := db.Prepare("DELETE FROM ressource WHERE id=?")
	help.CheckErr(err)

	reponse, err := stmt.Exec(id)
	help.CheckErr(err)

	affect, err := reponse.RowsAffected()
	help.CheckErr(err)

	if affect != 1 {
		return false
	}
	return true
}
