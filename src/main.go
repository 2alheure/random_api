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
			ENDPOINT /ping
	*/

	/**
	* @api {get} /ping
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiVersion 0.3.0
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("GET")
	/**
	* @api {post} /ping
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiVersion 0.3.0
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("POST")
	/**
	* @api {put} /ping
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiVersion 0.3.0
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("PUT")
	/**
	* @api {delete} /ping
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiVersion 0.3.0
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("DELETE")
	/**
	* @api {options} /ping
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", OptionsPing).Methods("OPTIONS")



	/**
			ENDPOINT /ressource
	*/

	/**
	* @api {get} /ressource?id={id}
	* @apiDescription Récupère la structure complète d'une ressource
	* @apiVersion 0.3.0
	* @apiGroup Ressource
	*
	* @apiParam {Number} id L'id de la ressource
	* @apiSuccess {Object} Ressource Les informations de la ressource
	* {
    *     "id": 1,
    *     "nom": "user",
    *     "createur": "2alheure",
    *     "date_creation": "2018-05-20 15:27:38",
    *     "champs": [
    *         {
    *             "id": 3,
    *             "clef": "age",
    *             "regle": {
    *                 "id": 11,
    *                 "nom": "Compris entre",
    *                 "parametres": [
    *                     {
    *                         "id": 13,
    *                         "type": "float",
    *                         "value": "0"
    *                     },
    *                     {
    *                         "id": 10,
    *                         "type": "float",
    *                         "value": "100"
    *                     }
    *                 ]
    *             }
    *         },
    *         {
    *             "id": 1,
    *             "clef": "nom",
    *             "regle": {
    *                 "id": 10,
    *                 "nom": "Dictionnaire",
    *                 "parametres": [
    *                     {
    *                         "id": 8,
    *                         "type": "string",
    *                         "value": "nom"
    *                     }
    *                 ]
    *             }
    *         },
    *         {
    *             "id": 2,
    *             "clef": "prenom",
    *             "regle": {
    *                 "id": 10,
    *                 "nom": "Dictionnaire",
    *                 "parametres": [
    *                     {
    *                         "id": 9,
    *                         "type": "string",
    *                         "value": "prenom"
    *                     }
    *                 ]
    *             }
    *         },
    *         {
    *             "id": 4,
    *             "clef": "sexe",
    *             "regle": {
    *                 "id": 1,
    *                 "nom": "Regex",
    *                 "parametres": [
    *                     {
    *                         "id": 11,
    *                         "type": "string",
    *                         "value": "[Homme|Femme]"
    *                     }
    *                 ]
    *             }
    *         },
    *         {
    *             "id": 5,
    *             "clef": "ville",
    *             "regle": {
    *                 "id": 10,
    *                 "nom": "Dictionnaire",
    *                 "parametres": [
    *                     {
    *                         "id": 12,
    *                         "type": "string",
    *                         "value": "ville"
    *                     }
    *                 ]
    *             }
    *         }
    *     ]
	* }
	*/
	router.HandleFunc("/ressource", controller.GetRessource).Methods("GET")
	/**
	* @api {post} /ressource
	* @apiDescription Crée une ressource
	* @apiVersion 0.3.0
	* @apiGroup Ressource
	*
	* @apiParam {String} nom Le nom de la ressource
	* @apiParam {String} createur Le nom du créateur de la ressource
	*
	* @apiSuccess (201 Created) {Object} Ressource Les informations de la ressource nouvellement créée
	*/
	router.HandleFunc("/ressource", controller.CreateRessource).Methods("POST")
	/**
	* @api {delete} /ressource
	* @apiDescription Supprime une ressource
	* @apiVersion 0.3.0
	* @apiGroup Ressource
	*
	* @apiParam {Number} id L'id de la ressource à supprimer
	*/
	router.HandleFunc("/ressource", controller.DeleteRessource).Methods("DELETE")
	/**
	* @api {options} /ressource
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Ressource
	*/
	router.HandleFunc("/ressource", controller.OptionsRessource).Methods("OPTIONS")



	/**
			ENDPOINT /ressources
	*/

	/**
	* @api {get} /ressources?max={max}
	* @apiDescription Récupère les informations minimales des ressources
	* @apiVersion 0.3.0
	* @apiGroup Ressources
	*
	* @apiParam {Number} max Le nombre maximum de ressources à récupérer
	* @apiSuccessExample {json} Exemple de retour :
	* [
	*     {
	*         "id": 1,
	*         "nom": "user",
	*         "createur": "2alheure",
	*         "date_creation": "2018-05-20 15:27:38"
	*     },
	*     {
	*         "id": 2,
	*         "nom": "machin",
	*         "createur": "frambur",
	*         "date_creation": "2018-05-20 17:34:43"
	*     }
	* ]
	*/
	router.HandleFunc("/ressources", controller.GetRessources).Methods("GET")
	/**
	* @api {options} /ressources
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Ressources
	*/
	router.HandleFunc("/ressources", controller.OptionsRessources).Methods("OPTIONS")



	/**
			ENDPOINT /champs
	*/

	/**
	* @api {get} /champs?max={max}
	* @apiDescription Récupère les informations minimales des champs
	* @apiVersion 0.3.0
	* @apiGroup Champs
	*
	* @apiParam {Number} max Le nombre maximum de champs à récupérer
	* @apiSuccessExample {json} Exemple de retour :
	* [
    *     {
    *         "id": 1,
    *         "clef": "nom"
    *     },
    *     {
    *         "id": 2,
    *         "clef": "prenom"
    *     },
    *     {
    *         "id": 3,
    *         "clef": "age"
    *     },
    *     {
    *         "id": 4,
    *         "clef": "sexe"
    *     },
    *     {
    *         "id": 5,
    *         "clef": "ville"
    *     }
	* ]
	*/
	router.HandleFunc("/champs", controller.GetChamps).Methods("GET")
	/**
	* @api {options} /champs
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Champs
	*/
	router.HandleFunc("/champs", controller.OptionsChamps).Methods("OPTIONS")



	/**
			ENDPOINT /champ
	*/

	/**
	* @api {get} /champ?id={id}
	* @apiDescription Récupère la structure complète d'un champ
	* @apiVersion 0.3.0
	* @apiGroup Champ
	*
	* @apiParam {Number} id L'id du champ
	* @apiSuccess {Object} Champ Les informations du champ nouvellement créé
	*/
	router.HandleFunc("/champ", controller.GetChamp).Methods("GET")
	/**
	* @api {post} /champ
	* @apiDescription Crée un champ
	* @apiVersion 0.3.0
	* @apiGroup Champ
	*
	* @apiParam {String} clef La clef du champ
	*
	* @apiSuccess (201 Created) {Object} Champ Les informations du champ nouvellement créé
	*/
	router.HandleFunc("/champ", controller.CreateChamp).Methods("POST")
	/**
	* @api {delete} /champ
	* @apiDescription Supprime un champ
	* @apiVersion 0.3.0
	* @apiGroup Champ
	*
	* @apiParam {Number} id L'id du champ à supprimer
	*/
	router.HandleFunc("/champ", controller.DeleteChamp).Methods("DELETE")
	/**
	* @api {options} /champ
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Champ
	*/
	router.HandleFunc("/champ", controller.OptionsChamp).Methods("OPTIONS")



	/**
			ENDPOINT /regles
	*/

	/**
	* @api {get} /regles?max={max}
	* @apiDescription Récupère les informations des règles
	* @apiVersion 0.3.0
	* @apiGroup Regles
	*
	* @apiParam {Number} max Le nombre maximum de regles à récupérer
	* @apiSuccessExample {json} Exemple de retour :
	* [
    *     {
    *         "id": 7,
    *         "nom": "Pair"
    *     },
    *     {
    *         "id": 1,
    *         "nom": "Regex",
    *         "parametres": [
    *             {
    *                 "id": 1,
    *                 "type": "string"
    *             }
    *         ]
    *     },
    *     {
    *         "id": 2,
    *         "nom": "Inférieur",
    *         "parametres": [
    *             {
    *                 "id": 3,
    *                 "type": "int"
    *             }
    *         ]
	*     }
	* ]
	*/
	router.HandleFunc("/regles", controller.GetRegles).Methods("GET")
	/**
	* @api {options} /regles
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiVersion 0.3.0
	* @apiGroup Regles
	*/
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
