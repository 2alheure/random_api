package controller

import (
	"net/http"
	"strconv"

	help "random_api/src/helper"
	model "random_api/src/model"
)

func GetChamp(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	champ := model.GetChamp(id)

	help.ReturnJson(w, champ)
}

func CreateChamp(w http.ResponseWriter, r *http.Request) {
	champ := model.Champ{Clef: r.FormValue("clef")}

	champ.Create()
	help.Return(w, 201, champ)
}

func DeleteChamp(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	help.CheckErr(err)

	champ := model.Champ{Id: int(id)}

	if champ.Delete() {
		help.ReturnOK(w)
	} else {
		help.Return(w, 404, `{"error": "A champ couldn't be deleted."}`)
	}
}

func OptionsChamp(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"POST",
		"DELETE",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}

