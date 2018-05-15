package helper

import (
	"net/http"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnJson(w http.ResponseWriter, js []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}