package generator;

import (
	"fmt"
	"github.com/lucasjones/reggen"
	help "random_api/src/helper"
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
	help.CheckErr(err)

	for i := 0; i < n; i++ {
		ret = append(ret, g.Generate(100));
	}

	return ret;
}