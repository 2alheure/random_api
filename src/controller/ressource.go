package controller;

import(
	model "../model"
	"net/http"
)

func CreateRessource(w http.ResponseWriter, r *http.Request) {
	model.CreateRessource("Test", "Florian");
}