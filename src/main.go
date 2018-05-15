package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "./controller"
	sql "./model"
	help "./helper"
)

func main() {
	sql.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/ping", Ping).Methods("GET")
	router.HandleFunc("/createRessource", controller.CreateRessource).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(`Pong`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	help.ReturnJson(w, js)
}
