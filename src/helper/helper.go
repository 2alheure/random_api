package helper

import (
	"net/http"
	json "encoding/json"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnJson(w http.ResponseWriter, json interface{}) {
	js, err := json.Marshal(json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
