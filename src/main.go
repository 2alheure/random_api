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
	router.HandleFunc("/ping", Ping).Methods("POST")
	router.HandleFunc("/ping", Ping).Methods("PUT")
	router.HandleFunc("/ping", Ping).Methods("DELETE")
	router.HandleFunc("/ping", OptionsPing).Methods("OPTIONS")

	router.HandleFunc("/ressource", controller.GetRessource).Methods("GET")
	router.HandleFunc("/ressource", controller.CreateRessource).Methods("POST")
	router.HandleFunc("/ressource", controller.DeleteRessource).Methods("DELETE")
	router.HandleFunc("/ressource", controller.OptionsRessource).Methods("OPTIONS")

	/**
	* @api {get} /ressources?max=:max Request to get ressources
	* @apiVersion 1.0.0
	* @apiName GetRessources
	* @apiGroup Ressources
	*
	* @apiParam {Number} max Set how many ressources you want to get
	*/
	router.HandleFunc("/ressources", controller.GetRessources).Methods("GET")
	router.HandleFunc("/ressources", controller.OptionsRessources).Methods("OPTIONS")

	router.HandleFunc("/champs", controller.GetChamps).Methods("GET")
	router.HandleFunc("/champs", controller.OptionsChamps).Methods("OPTIONS")
	
	router.HandleFunc("/champ", controller.GetChamp).Methods("GET")
	router.HandleFunc("/champ", controller.CreateChamp).Methods("POST")
	router.HandleFunc("/champ", controller.DeleteChamp).Methods("DELETE")
	router.HandleFunc("/champ", controller.OptionsChamp).Methods("OPTIONS")

	router.HandleFunc("/regles", controller.GetRegles).Methods("GET")
	router.HandleFunc("/regles", controller.OptionsRegles).Methods("OPTIONS")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	help.ReturnJson(w, `Pong`)
}

func OptionsPing(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
