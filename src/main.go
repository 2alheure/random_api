package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "random_api/src/controller"
	help "random_api/src/helper"
	model "random_api/src/model"
	config "random_api/src/config"
)

func main() {
	model.InitBdd()

	defer model.Bdd.Close()

	router := mux.NewRouter()

	/**
	* @api {get} / Home
	* @apiDescription Renvoie le lien vers la <a href="http://2dtension.fr/alea-data-est">documentation de l'API</a>
	* @apiGroup Home
	* @apiSuccessExample {json} Retourne :
	* {"documentation": "http://2dtension.fr/alea-data-est"}
	*/
	router.HandleFunc("/", Home).Methods("GET")
	/**
	* @api {options} / Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Home
	*/
	router.HandleFunc("/", OptionsHome).Methods("OPTIONS")


	/**
				ENDPOINT /ping
	*/
{
	
	/**
	* @api {get} /ping Get
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("GET")
	/**
	* @api {post} /ping Post
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("POST")
	/**
	* @api {put} /ping Put
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("PUT")
	/**
	* @api {delete} /ping Delete
	* @apiDescription Ping le serveur pour vérifier qu'il fonctionne
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", Ping).Methods("DELETE")
	/**
	* @api {options} /ping Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Ping
	*/
	router.HandleFunc("/ping", OptionsPing).Methods("OPTIONS")
}


	/**
				ENDPOINT /ressource
	*/
{
	/**
	* @api {get} /ressource?id={id} Structure
	* @apiDescription Récupère la structure complète d'une ressource
	* @apiGroup Ressource
	*
	* @apiParam {Number} id L'id de la ressource
	* @apiSuccess {Object} Ressource Les informations de la ressource
	* @apiSuccessExample {json} Exemple de retour :
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
	* @api {post} /ressource Créer
	* @apiDescription Crée une ressource
	* @apiGroup Ressource
	*
	* @apiParam {String} nom Le nom de la ressource
	* @apiParam {String} createur Le nom du créateur de la ressource
	*
	* @apiSuccess (Created 201) {Object} Ressource Les informations de la ressource nouvellement créée
	* @apiSuccessExample {json} Exemple de retour :
	* {
	* 	"id": 6,
	* 	"nom": "ressource",
	* 	"createur": "frambur",
	* 	"date_creation": "2018-05-28 22:52:26"
	* }
	*/
	router.HandleFunc("/ressource", controller.CreateRessource).Methods("POST")
	/**
	* @api {put} /ressource Modifier
	* @apiDescription Modifie une ressource
	* @apiGroup Ressource
	*
	* @apiParam {String} id L'id de la ressource à modifier
	* @apiParam {String} [nom] Le nouveau nom de la ressource
	* @apiParam {Number} [createur] Le nom de la personne à qui réaffecter la création de la ressource
	*/
	router.HandleFunc("/ressource", controller.ModifyRessource).Methods("PUT")
	/**
	* @api {delete} /ressource Supprimer
	* @apiDescription Supprime une ressource
	* @apiGroup Ressource
	*
	* @apiParam {Number} id L'id de la ressource à supprimer
	*/
	router.HandleFunc("/ressource", controller.DeleteRessource).Methods("DELETE")
	/**
	* @api {options} /ressource Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Ressource
	*/
	router.HandleFunc("/ressource", controller.OptionsRessource).Methods("OPTIONS")
}


	/**
				ENDPOINT /ressources
	*/
{
	/**
	* @api {get} /ressources?max={max} Lister
	* @apiDescription Récupère les informations minimales des ressources
	* @apiGroup Ressources
	*
	* @apiParam {Number} max Le nombre maximum de ressources à récupérer
	* @apiSuccess {Object} Regle[] Les informations minimales des ressources
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
	* @api {options} /ressources Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Ressources
	*/
	router.HandleFunc("/ressources", controller.OptionsRessources).Methods("OPTIONS")
}


	/**
				ENDPOINT /champs
	*/
{
	/**
	* @api {get} /champs?max={max} Lister
	* @apiDescription Récupère les informations minimales des champs
	* @apiGroup Champs
	*
	* @apiParam {Number} max Le nombre maximum de champs à récupérer
	* @apiSuccess {Object} Champ[] Les informations minimales des champs
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
	* @api {options} /champs Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Champs
	*/
	router.HandleFunc("/champs", controller.OptionsChamps).Methods("OPTIONS")
}


	/**
				ENDPOINT /champ
	*/
{
	/**
	* @api {get} /champ?id={id} Structure
	* @apiDescription Récupère la structure complète d'un champ
	* @apiGroup Champ
	*
	* @apiParam {Number} id L'id du champ
	* @apiSuccess {Object} Champ Les informations du champ
	* @apiSuccessExample {json} Exemple de retour :
	* {
	* 	"id": 9,
	* 	"clef": "machin",
	* 	"ressource_id": null,
	* 	"regle": {
	* 		"id": 11,
	* 		"nom": "Compris entre",
	* 		"parametres": [
	* 			{
	* 				"id": 4,
	* 				"type": "float"
	* 			},
	* 			{
	* 				"id": 4,
	* 				"type": "float"
	* 			}
	* 		]
	* 	}
	* }
	*/
	router.HandleFunc("/champ", controller.GetChamp).Methods("GET")
	/**
	* @api {post} /champ Créer
	* @apiDescription Crée un champ
	* @apiGroup Champ
	*
	* @apiParam {String} clef La clef du champ
	* @apiParam {Number} [ressource_id] L'id de la ressource à laquelle le champ se rattache
	*
	* @apiSuccess (Created 201) {Object} Champ Les informations du champ nouvellement créé
	* @apiSuccessExample {json} Exemple de retour :
	* {
	* 	"id": 10,
	* 	"clef": "champ",
	* 	"ressource_id": null
	* }	
	*/
	router.HandleFunc("/champ", controller.CreateChamp).Methods("POST")
	/**
	* @api {put} /champ Modifier
	* @apiDescription Modifie un champ
	* @apiGroup Champ
	*
	* @apiParam {String} id L'id du champ à modifier
	* @apiParam {String} [clef] La nouvelle clef du champ
	* @apiParam {Number} [ressource_id] Le nouvel id de la ressource à laquelle le champ se rattache <br /><br />
	* Si vous souhaitez ne plus attacher le champ à une ressource, renseignez la clef et laissez sa valeur vide.
	*/
	router.HandleFunc("/champ", controller.ModifyChamp).Methods("PUT")
	/**
	* @api {delete} /champ Supprimer
	* @apiDescription Supprime un champ
	* @apiGroup Champ
	*
	* @apiParam {Number} id L'id du champ à supprimer
	*/
	router.HandleFunc("/champ", controller.DeleteChamp).Methods("DELETE")
	/**
	* @api {options} /champ Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Champ
	*/
	router.HandleFunc("/champ", controller.OptionsChamp).Methods("OPTIONS")
}


	/**
				ENDPOINT /regles
	*/
{
	/**
	* @api {get} /regles?max={max} Lister
	* @apiDescription Récupère les informations des règles
	* @apiGroup Regles
	*
	* @apiParam {Number} max Le nombre maximum de regles à récupérer
	* @apiSuccess {Object} Regle[] Les informations des règles
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
	* @api {options} /regles Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Regles
	*/
	router.HandleFunc("/regles", controller.OptionsRegles).Methods("OPTIONS")
}

	/**
				ENDPOINT /regle
	*/
{
	/**
	* @api {get} /regle?id={id} Structure
	* @apiDescription Récupère les informations des règles
	* @apiGroup Regle
	*
	* @apiParam {Number} id L'id de la règle à récupérer
	* @apiSuccess {Object} Regle Les informations de al règle
	* @apiSuccessExample {json} Exemple de retour :
	* {
    *     "id": 2,
    *     "nom": "Inférieur",
    *     "parametres": [
    *         {
    *             "id": 3,
    *             "type": "int"
    *         }
    *     ]
	* }
	* 
	*/
	router.HandleFunc("/regle", controller.GetRegle).Methods("GET")
	/**
	* @api {post} /regle Assigner
	* @apiDescription Assigne une règle à un champ
	* @apiGroup Regle
	*
	* @apiParam {String} regle_id L'id de la règle à assigner
	* @apiParam {String} champ_id L'id du champ auquel sera assignée la règle
	*/
	router.HandleFunc("/regle", controller.AttachRegle).Methods("POST")
	/**
	* @api {delete} /regle Désassigner
	* @apiDescription Désssigne sa règle à un champ
	* @apiGroup Regle
	*
	* @apiParam {String} champ_id L'id du champ duquel désassigner sa règle
	*/
	router.HandleFunc("/regle", controller.DetachRegle).Methods("DELETE")
	/**
	* @api {options} /regle Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Regle
	*/
	router.HandleFunc("/regle", controller.OptionsRegle).Methods("OPTIONS")
}

	/**
				ENDPOINT /parametres
	*/
{
	/**
	* @api {post} /parametres (Re)définir
	* @apiDescription (Re)définit les paramètres de la règle d'un champ donné
	* @apiGroup Parametres
	*
	* @apiParam {String} champ_id L'id du champ duquel les paramètres seront (re)définis
	* @apiParam {String[]} parametres Les paramètres <br /><br />
	* Les paramètres sont à mettre, en json, dans le corps de la requête.<br /><br />
	* Ils doivent être sous forme de tableau de chaînes de caractères.
	* @apiParamExample {json} Exemple d'envoi de paramètres :
	* [
	* 	"1.3",
	* 	"true",
	* 	"42",
	* 	"Hello, World"
	* ]
	*/
	router.HandleFunc("/parametres", controller.SetParametres).Methods("POST")
	/**
	* @api {put} /parametres (Re)définir
	* @apiDescription (Re)définit les paramètres de la règle d'un champ donné
	* @apiGroup Parametres
	*
	* @apiParam {String} champ_id L'id du champ duquel les paramètres seront (re)définis
	* @apiParam {String[]} parametres Les paramètres <br /><br />
	* Les paramètres sont à mettre, en json, dans le corps de la requête.<br /><br />
	* Ils doivent être sous forme de tableau de chaînes de caractères.
	* @apiParamExample {json} Exemple d'envoi de paramètres :
	* [
	* 	"1.3",
	* 	"true",
	* 	"42",
	* 	"Hello, World"
	* ]
	*/
	router.HandleFunc("/parametres", controller.SetParametres).Methods("PUT")
	/**
	* @api {delete} /parametres Réinitialiser
	* @apiDescription Réinitialise les paramètres d'une règle d'un champ et remet leur valeur à <code>null</code>
	* @apiGroup Parametres
	*
	* @apiParam {String} champ_id L'id du champ duquel les paramètres seront réinitialisés
	*/
	router.HandleFunc("/parametres", controller.ResetParametres).Methods("DELETE")
	/**
	* @api {options} /parametres Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Parametres
	*/
	router.HandleFunc("/parametres", controller.OptionsParametres).Methods("OPTIONS")
}

	/**
				ENDPOINT /generate
	*/
{
	/**
	* @api {get} /generate?ressource_id={ressource_id}&nombre={nombre} Générer
	* @apiDescription Génère aléatoirement une ou plusieurs ressources
	* @apiGroup Generate
	*
	* @apiParam {Number} ressource_id L'id de la ressource à générer
	* @apiParam {Number} [nombre] Le nombre d'instances à renvoyer <br /><br />
	* Doit être un nombre entier et positif.<br /><br />
	* Si non renseigné, renvoie une seule instance.
	* @apiSuccess {[]Ressource} Ressource Un tableau de {nombre} ressource
	* @apiSuccessExample {json} Exemple de retour :
	* [
	* 	"1.3",
	* 	"true",
	* 	"42",
	* 	"Hello, World"
	* ]
	*/
	router.HandleFunc("/generate", controller.Generate).Methods("GET")
	/**
	* @api {options} /generate Options
	* @apiDescription Renvoie la liste des méthodes autorisées
	* @apiGroup Generate
	*/
	router.HandleFunc("/generate", controller.OptionsGenerate).Methods("OPTIONS")
}



	log.Fatal(http.ListenAndServe(config.Port, router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	help.ReturnJson(w, `"Pong"`)
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

func Home(w http.ResponseWriter, r *http.Request) {
	help.ReturnJson(w, `{"documentation": "http://2dtension.fr/alea-data-est"}`)
}

func OptionsHome(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
