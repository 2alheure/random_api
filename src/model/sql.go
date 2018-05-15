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

/**
func GetRessources(max int) *[]struct {} {
	var res []struct {
		nom string
		createur string
	}

	stmt, err := db.Query("SELECT * FROM ressource LIMIT ?")
	if err != nil {
		panic(err.Error())
	}

	for stmt.Next() {
		var nom string
		var createur string
		err = stmt.Scan(&nom, &createur)
		if err != nil {
			panic(err.Error())
		}

		
		res = append(res, struct {
			nom string
			createur string
		} { nom, createur })
	}

	return &res
}
*/

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
