package helper

import (
	"fmt"
	"net/http"
	"strings"
	json "encoding/json"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ReturnJson(w http.ResponseWriter, to_json interface{}) {
	js, err := json.Marshal(to_json)
	if err != nil {
		http.Error(w, "An Error Has Occured", http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ReturnOptions(w http.ResponseWriter, options []string) {
	opt := strings.Join(options, ", ")
	w.Header().Set("Access-Control-Allow-Methods", opt)
	w.Write(nil)
}

func Return(w http.ResponseWriter, code int, to_json interface{}) {
	if to_json != nil{
		w.Header().Set("Content-Type", "application/json")	
		w.WriteHeader(code)
		ReturnJson(w, to_json)
	} else {
		w.WriteHeader(code)
	}
}

func ReturnOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
