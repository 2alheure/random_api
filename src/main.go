package main

import (
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

	controller "./controller"
	help "./helper"
)

func main() {
	db, err := sql.Open("mysql", "root:@/alea_data_est?charset=utf8")
	help.CheckErr(err)

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/ping", Ping).Methods("GET")

	router.HandleFunc("/ressource", controller.GetRessource).Methods("GET")
	router.HandleFunc("/ressource", controller.CreateRessource).Methods("POST")
	router.HandleFunc("/ressource", controller.DeleteRessource).Methods("DELETE")
	router.HandleFunc("/ressources", controller.GetRessources).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	help.ReturnJson(w, `Pong`)
}
