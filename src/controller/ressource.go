package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	model "../model"
	help "../helper"
)

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	nom := r.FormValue("nom");
	createur := r.FormValue("createur");
	model.CreateRessource(nom, createur);
	fmt.Println("OK");
	js, err := json.Marshal(`{OK}`);
	help.CheckErr(err);
	help.ReturnJson(w, js);
}

func GetRessources(w http.ResponseWriter, r *http.Request) {
	max := r.FormValue("max");

	results, err := model.GetRessources(max);
	help.CheckErr(err);

	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	help.ReturnJson(w, js);
}

func DeleteRessource(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"));
	help.CheckErr(err);
	model.DeleteRessource(id);
	fmt.Println("OK");
	js, err := json.Marshal(`{OK}`);
	help.CheckErr(err);
	help.ReturnJson(w, js);
}