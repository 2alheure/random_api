package controller

import (
	"net/http"
	"strconv"

	help "random_api/src/helper"
	model "random_api/src/model"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	ressource_id := r.FormValue("ressource_id")

	if ressource_id != "" {
		if nombre != "" {
			nombre, err := strconv.Atoi(nombre)
		} else {
			nombre = 1
		}

		ressource_id, err2 := strconv.Atoi(ressource_id)

		if err != nil || err2 != nil {
			help.Return(w, 400, nil)
		} else {
			max := 25			// Première mesure anti-DDOS, un max

			if nombre <= max {
				ret, err_code := model.Generate(ressource_id, nombre)
	
				if err_code != 0 {
					help.Return(w, err_code, nil)
				} else {
					help.ReturnJson(w, ret)
				}
			} else {
				help.Return(w, 413, nil)
			}
		}
	} else {
		help.Return(w, 400, nil)
	}
}