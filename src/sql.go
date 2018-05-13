package model;

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
);

func checkErr(err error) {
	if (err != nil) {
		panic(err);
	}
}

func connect() {
	db, err := sql.Open("mysql", "root:@/alea_data_est?charset=utf8");
	checkErr(err);
	return db;
}

func CreateRessource(nom, createur string) {
	db = connect();
	defer db.Close();

	stmt, err := db.Prepare("INSERT INTO ressource (nom, createur) VALUE (?, ?)");
	checkErr(err);

	reponse, err := stmt.Exec(nom, createur);
	checkErr(err);

	id, err := res.LastInsertId();
	checkErr(err);

	return id;
}

func DeleteRessource(id int) bool {
	db = connect();
	defer db.Close();

	stmt, err = db.Prepare("DELETE FROM ressource WHERE id=?");
	checkErr(err);

	reponse, err = stmt.Exec(id);
	checkErr(err);

	affect, err = reponse.RowsAffected();
	checkErr(err);

	if (affect != 1) return false;
	return true;
}