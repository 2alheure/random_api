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
