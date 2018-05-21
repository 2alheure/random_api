package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "./controller"
	help "./helper"
	model "./model"
)

func main() {
	Bdd := model.InitBdd()

	defer Bdd.Close()

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
