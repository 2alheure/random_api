package main;

import (
	"encoding/json"
	"github.com/gorilla/mux"
    "log"
	"net/http"
	controller "./controller"
)

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/ping", Ping).Methods("GET");
	router.HandleFunc("/createRessource", controller.CreateRessource).Methods("POST");
    log.Fatal(http.ListenAndServe(":8000", router));
}

func Ping(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal("Pong")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError);
		panic(err);
	}
	
	w.Header().Set("Content-Type", "application/json");
	w.Write(js);
}