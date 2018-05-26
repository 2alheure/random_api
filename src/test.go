package main;

import (
    "fmt"
);

func main() {
    fmt.Printf("hello, world\n");
    test := generator.FromRegex("test", 10);
    fmt.Println(generator.LowerThan(5));
    fmt.Println(generator.StrictlyLowerThan(5));

    fmt.Println(test);
}