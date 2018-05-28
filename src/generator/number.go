package generator;

import (
	"fmt"
	rand "math/rand"
	"time"
);

func TestNum() {
	fmt.Println("test num_gen");
}

func BetweenMinAndMax(min int, max int) int {
	rand.Seed(time.Now().UnixNano());
	var ret int;
	ret = rand.Intn(max - min) + min;

	return ret;
}

func LowerThan(n int) int {
	rand.Seed(time.Now().UnixNano());
	var ret int;
	ret = rand.Intn(n+1);

	return ret;
}

func StrictlyLowerThan(n int) int {
	rand.Seed(time.Now().UnixNano());
	var ret int;
	ret = rand.Intn(n);

	return ret;
}

/**
func GreaterThan(n int) int {
	rand.Seed(time.Now().UnixNano());
	var ret int;
	ret = rand.Int();

	return ret;
}

func strictlyGreaterThan() int {

}
*/