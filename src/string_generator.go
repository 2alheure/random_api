package generator;

import (
	"fmt"
	"github.com/lucasjones/reggen"
);

func FromRegex(regex string, n int) []string {
	if (n <= 0) return [];

	var ret [n]string; 

	g, err := reggen.NewGenerator(regex);
	if err != nil {
		panic(err);
	}

	for (i := 0; i < n; i++) {
		ret[i] = g.Generate(100);
	}

	return ret;
}