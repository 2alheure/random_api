package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "random_api/src/controller"
	help "random_api/src/helper"
	model "random_api/src/model"
)

func main() {
	model.InitBdd()

	defer model.Bdd.Close()

	router := mux.NewRouter()

	/**
	* @api {get} /ping Request to ping the server and check if it's working
	* @apiVersion 1.0.0
	* @apiName Ping
	* @apiGroup Ping
	* @apiSuccess {String} pong The user pong
	*/
	router.HandleFunc("/ping", Ping).Methods("GET")

	router.HandleFunc("/ressource", controller.GetRessource).Methods("GET")
	router.HandleFunc("/ressource", controller.CreateRessource).Methods("POST")
	router.HandleFunc("/ressource", controller.DeleteRessource).Methods("DELETE")
	
	/**
	* @api {get} /ressource/:max Request to get ressources
	* @apiVersion 1.0.0
	* @apiName GetRessources
	* @apiGroup Ressources
	*
	* @apiParam {Number} max Set how many ressources you want to get
	*/
	router.HandleFunc("/ressources", controller.GetRessources).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	help.ReturnJson(w, `Pong`)
}