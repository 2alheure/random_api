package controller

import (
	"net/http"
	"strconv"

	help "../helper"
	model "../model"
)

func GetChamps(w http.ResponseWriter, r *http.Request) {
	max, err := strconv.Atoi(r.FormValue("max"))
	help.CheckErr(err)

	results := model.GetChamps(max)

	help.ReturnJson(w, results)
}

func OptionsChamps(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
