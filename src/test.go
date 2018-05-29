package main;

import (
    "fmt"
    "encoding/json"
    model "random_api/src/model"
);

func main() {
	model.InitBdd()

    defer model.Bdd.Close()
    

    reducer, err_code, erroring := model.GetReducer(7)

    fmt.Println(reducer)
    fmt.Println(err_code)
    fmt.Println(erroring)



    ress, code := model.Generate(1, 10)

    fmt.Println(code)

    if (code == 200) {
        fmt.Println(json.Marshal(ress))
    }
}