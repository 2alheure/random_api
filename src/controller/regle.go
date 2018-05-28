package controller

import (
	"net/http"
	"strconv"

	help "random_api/src/helper"
	model "random_api/src/model"
)

func GetRegle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	results := model.GetRegle(id)

	help.ReturnJson(w, results)
}

func AttachRegle(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")
	regle_id := r.FormValue("regle_id")

	if champ_id == "" || regle_id == "" {
		help.Return(w, 400, nil)
	} else {
		champ_id, err := strconv.Atoi(champ_id)
		help.CheckErr(err)
		
		regle_id, err := strconv.Atoi(regle_id)
		help.CheckErr(err)

		if model.AttachRule(regle_id, champ_id) {
			help.ReturnOK(w)
		} else {
			help.Return(w, 404, nil)
		}
	}
}

func DetachRegle(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")

	if champ_id == "" {
		help.Return(w, 400, nil)
	} else {
		champ_id, err := strconv.Atoi(champ_id)
		help.CheckErr(err)

		if model.DetachRule(champ_id) {
			help.ReturnOK(w)
		} else {
			help.Return(w, 404, nil)
		}
	}
}

func OptionsRegle(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"POST",
		"DELETE",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
