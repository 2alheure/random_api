package generator;

import (
	"fmt"
	rand "math/rand"
	constante "math/const"
	"errors"
	"strconv"
	"time"

	model "../src/model"
);

func TestNum() {
	fmt.Println("test num_gen");
}

func BetweenMinAndMax(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 2 {
		minStr := args[0]
		maxStr := args[1]

		min, err := strconv.Atoi(minStr)
		if err != nil {
			return toReturn, err
		}

		max, err := strconv.Atoi(maxStr)
		if err != nil {
			return toReturn, err
		}

		if min > max {
			return BetweenMinAndMax(maxStr, minStr)
		}

		rand.Seed(time.Now().UnixNano());
		var ret int;
		ret = rand.Intn(max - min) + min;

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func Equal(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		nStr := args[0]

		n, err := strconv.Atoi(nStr)
		if err != nil {
			return toReturn, err
		}

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func LowerThan(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		nStr := args[0]
		
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return toReturn, err
		}

		rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(n+1)

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func StrictlyLowerThan(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		nStr := args[0]

		n, err := strconv.Atoi(nStr)
		if err != nil {
			return toReturn, err
		}

		rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(n)

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func EvenNumber(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 0 {
		rand.Seed(time.Now().UnixNano())
		var ret int
		ret = (rand.Intn(int(constante.MaxInt32))) * 2

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func OddNumber(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 0 {
		rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(int(constante.MaxInt32))
		ret += (ret % 2) + 1

		return ChampKV{Int: ret}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

