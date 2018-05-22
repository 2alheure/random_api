package controller

import (
	"net/http"
	"strconv"

	help "random_api/src/helper"
	model "random_api/src/model"
)

func GetRessources(w http.ResponseWriter, r *http.Request) {
	max, err := strconv.Atoi(r.FormValue("max"))
	help.CheckErr(err)

	results := model.GetRessources(max)

	help.ReturnJson(w, results)
}

func GetRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	ress := model.GetRessource(id)

	help.ReturnJson(w, ress)
}

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	ress := model.Ressource{Nom: r.FormValue("nom"), Createur: r.FormValue("createur")}

	ress.Create()
	help.ReturnJson(w, ress)
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
