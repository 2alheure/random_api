package main;

import (
    "fmt"
    // "encoding/json"
    model "./model"
);

func main() {
	model.InitBdd()

    defer model.Bdd.Close()
    

    reducer, err_code, erroring := model.GetReducer(7)

    fmt.Println(reducer)
    fmt.Println(err_code)
    fmt.Println(erroring)


    // ress, code := model.Generate(7, 10)

    // fmt.Println(code)

    // if (code == 200) {
    //     js, err :=json.Marshal(ress)
        
    //     fmt.Println(err)
    //     fmt.Println(string(js))
    // }
}