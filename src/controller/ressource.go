package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "../model"
	help "../helper"
)

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	model.CreateRessource("Test", "Florian")
	fmt.Println("OK")
	js, err := json.Marshal(`{OK}`)
	help.CheckErr(err)
	help.ReturnJson(w, js)
}
