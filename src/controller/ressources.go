package controller

import (
	"net/http"
	"strconv"

	help "../helper"
	model "../model"
)

func GetRessources(w http.ResponseWriter, r *http.Request) {
	max, err := strconv.Atoi(r.FormValue("max"))
	help.CheckErr(err)

	results := model.GetRessources(max)

	help.ReturnJson(w, results)
}

func OptionsRessources(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
