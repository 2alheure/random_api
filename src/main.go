package main;

import (
	"fmt"
	"./generator"
	"./model"
)

func main() {
	fmt.Println("test");
	generator.TestNum();
	generator.TestStr();
	model.TestSql();
}