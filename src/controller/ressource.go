package controller

import (
	"net/http"
	"strconv"

	help "random_api/src/helper"
	model "random_api/src/model"
)

func GetRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	ress := model.GetRessource(id)
	ress.Hydrate()

	help.ReturnJson(w, ress)
}

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	ress := model.Ressource{Nom: r.FormValue("nom"), Createur: r.FormValue("createur")}

	ress.Create()
	help.Return(w, 201, ress)
}

func ModifyRessource(w http.ResponseWriter, r *http.Request) {
	ress := model.Ressource{Nom: r.FormValue("nom"), Createur: r.FormValue("createur")}

	ress.Create()
	help.Return(w, 201, ress)
}

func DeleteRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	ress := model.Ressource{Id: int(id)}

	if ress.Delete() {
		help.ReturnJson(w, `{OK}`)
	} else {
		help.ReturnJson(w, `{error: "A ressource couldn't be deleted."}`)
	}
}

func OptionsRessource(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"POST",
		"DELETE",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}

