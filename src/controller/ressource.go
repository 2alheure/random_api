package controller

import (
	"net/http"
	"strconv"
    "gopkg.in/guregu/null.v3"

	help "../helper"
	model "../model"
)

func GetRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	ress, err := model.GetRessource(id)

	if err != nil {
		help.Return(w, 404, nil)
	} else {
		ress.Hydrate()
		help.ReturnJson(w, ress)
	}
}

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	ress := model.Ressource{Nom: null.NewString(r.FormValue("nom"), true), Createur: null.NewString(r.FormValue("createur"), true)}

	ress.Create()
	help.Return(w, 201, ress)
}

func ModifyRessource(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		help.Return(w, 400, nil)
	} else {
		id, err := strconv.Atoi(r.FormValue("id"))
		help.CheckErr(err)
	
		ress, err := model.GetRessource(id);
		
		if err != nil {
			help.Return(w, 404, nil)
		} else {
			err = r.ParseForm()
			help.CheckErr(err)
			
			ress.Modify(r.Form)
			help.ReturnOK(w)
		}
	}
}

func DeleteRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	ress := model.Ressource{Id: int(id)}

	if ress.Delete() {
		help.ReturnOK(w)
	} else {
		help.Return(w, 404, nil)
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

