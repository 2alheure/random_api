package controller

import (
	"net/http"
	"strconv"

	help "../helper"
	model "../model"
)

func Generate(w http.ResponseWriter, r *http.Request) {
	nombre := r.FormValue("nombre")
	ressource_id := r.FormValue("ressource_id")
	var nbr int

	if ressource_id != "" {
		if nombre != "" {
			var err error
			nbr, err = strconv.Atoi(nombre)
			if err != nil {
				help.Return(w, 400, nil)
			}
		} else {
			nbr = 1
		}

		ressource_id, err2 := strconv.Atoi(ressource_id)

		if err2 != nil {
			help.Return(w, 400, nil)
		} else {
			max := 25			// Première mesure anti-DDOS, un max

			if nbr <= max {
				ret, err_code := model.Generate(ressource_id, nbr)
	
				if err_code != 200 {
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

func TestGenerate(w http.ResponseWriter, r *http.Request) {
	ressource_id := r.FormValue("ressource_id")

	var test struct {
		Status			string		`json:"status"`
		ErrorCode		int			`json:"error_code,omitempty"`
		ErrorMessages	[]string	`json:"debug_messages,omitempty"`
	}

	if ressource_id == "" {
		
		test.Status = "error"
		test.ErrorCode = 400
		test.ErrorMessages = append(test.ErrorMessages, "Le paramètre {id} est requis.")
		help.ReturnJson(w, test)
	} else {
		
		id, err := strconv.Atoi(ressource_id)
		if err != nil {
			
			test.Status = "error"
			test.ErrorCode = 400
			test.ErrorMessages = append(test.ErrorMessages, "Le paramètre {id} doit être un nombre entier.")
			help.ReturnJson(w, test)
		} else {
			
			errCode, messages := model.TestGenerate(id)
			if errCode != 200 {
				test.Status = "error"
				test.ErrorCode = errCode
				test.ErrorMessages = messages
				help.ReturnJson(w, test)	
			} else {
				test.Status = "success"
				help.ReturnJson(w, test)	
			}
		}
	}
}

func OptionsTest(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}

func OptionsGenerate(w http.ResponseWriter, r *http.Request) {
	options := []string{
		"GET",
		"OPTIONS",
	}

	help.ReturnOptions(w, options)
}
