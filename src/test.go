package main;

import (
    "fmt"
);

func main() {
    fmt.Printf("hello, world\n");
    test := generator.FromRegex("test", 10);

    fmt.Println(test);
}