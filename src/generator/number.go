package generator;

import (
	"fmt"
	rand "math/rand"
);

func TestNum() {
	fmt.Println("test num_gen");
}

func LowerThan(n int) int {
	var ret int;
	ret = rand.Intn(n+1);

	return ret;
}

func StrictlyLowerThan(n int) int {
	var ret int;
	ret = rand.Intn(n);

	return ret;
}

/**
func greaterThan(n int) int {

}

func strictlyGreaterThan() int {

}
*/