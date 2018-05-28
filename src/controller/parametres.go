package controller

import (
	"net/http"
	// "strconv"
	"fmt"
	json "encoding/json"

	help "random_api/src/helper"
	// model "random_api/src/model"
)

func SetParametres(w http.ResponseWriter, r *http.Request) {
	// champ_id := r.FormValue("champ_id")
	parametres := r.FormValue("parametres")

	var p []string
	err := json.Unmarshal([]byte(parametres), &p)

	help.CheckErr(err)

	fmt.Println(p)
}

func ResetParametres(w http.ResponseWriter, r *http.Request) {
	// champ_id := r.FormValue("champ_id")
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
