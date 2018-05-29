package controller

import (
	"net/http"
	"strconv"
	_ "fmt"
	json "encoding/json"

	help "../helper"
	model "../model"
)

func SetParametres(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")
	parametres := r.FormValue("parametres")

	var params []string
	err := json.Unmarshal([]byte(parametres), &params)
	
	if champ_id == "" || err != nil {
		help.Return(w, 400, nil)
	} else {
		champ_id, err := strconv.Atoi(champ_id)
		help.CheckErr(err)
		if model.SetParametres(champ_id, params) {
			help.ReturnOK(w)
		} else {
			help.Return(w, 400, nil)
		}
	}
}

func ResetParametres(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")

	if champ_id == "" {
		help.Return(w, 400, nil)
	} else {
		champ_id, err := strconv.Atoi(champ_id)
		help.CheckErr(err)
		if model.ResetParametres(champ_id) {
			help.ReturnOK(w)
		} else {
			help.Return(w, 404, nil)
		}
	}
}

func OptionsParametres(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"POST",
		"PUT",
		"DELETE",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
