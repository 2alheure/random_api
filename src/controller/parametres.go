package controller

import (
	"net/http"
	// "strconv"
	_ "fmt"
	json "encoding/json"

	help "random_api/src/helper"
	// model "random_api/src/model"
)

func SetParametres(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")
	parametres := r.FormValue("parametres")

	var params []string
	err := json.Unmarshal([]byte(parametres), &params)
	
	if champ_id == "" || err != nil {
		Return(w, 400, nil)
	} else {
		if model.SetParametres(champ_id, parametres) {
			ReturnOk(w)
		} else {
			Return(w, 404, nil)
		}
	}
}

func ResetParametres(w http.ResponseWriter, r *http.Request) {
	champ_id := r.FormValue("champ_id")

	if champ_id == "" {
		Return(w, 400, nil)
	} else {
		if model.ResetParametres(champ_id) {
			ReturnOk(w)
		} else {
			Return(w, 404, nil)
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
