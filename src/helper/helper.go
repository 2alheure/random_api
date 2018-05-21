package helper

import (
	"log"
	"net/http"
	json "encoding/json"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ReturnJson(w http.ResponseWriter, to_json interface{}) {
	js, err := json.Marshal(to_json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
