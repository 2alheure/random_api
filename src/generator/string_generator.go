package generator;

import (
	"fmt"
	"github.com/lucasjones/reggen"
);

func TestStr() {
	fmt.Println("test str_gen");
}

func FromRegex(regex string, n int) []string {
	if (n <= 0) {
		return nil;
	}

	var ret []string; 

	g, err := reggen.NewGenerator(regex);
	if (err != nil) {
		panic(err);
	}

	for i := 0; i < n; i++ {
		ret = append(ret, g.Generate(100));
	}

	return ret;
}