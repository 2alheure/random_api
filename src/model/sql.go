package model;

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
);

var db *sql.DB;

func TestSql() {
	fmt.Println("test sql");
}

func checkErr(err error) {
	if (err != nil) {
		panic(err);
	}
}

func connect() *sql.DB {
	var err error;
	db, err = sql.Open("mysql", "root:@/alea_data_est?charset=utf8");
	checkErr(err);
	return db;
}

func CreateRessource(nom, createur string) int64 {
	stmt, err := db.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)");
	checkErr(err);

	reponse, err := stmt.Exec(nom, createur);
	checkErr(err);

	id, err := reponse.LastInsertId();
	checkErr(err);

	return id;
}

func DeleteRessource(id int) bool {
	stmt, err := db.Prepare("DELETE FROM ressource WHERE id=?");
	checkErr(err);

	reponse, err := stmt.Exec(id);
	checkErr(err);

	affect, err := reponse.RowsAffected();
	checkErr(err);

	if (affect != 1) {
		return false;
	}
	return true;
}