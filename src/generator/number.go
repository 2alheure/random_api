package generator

import (
	"fmt"
	"math"
	"math/rand"
	"errors"
	"strconv"
	cryptrand "crypto/rand"
	"encoding/binary"
	"log"
)

func TestNum() {
	fmt.Println("test num_gen")
}

func resetSeed() {
	var seed int64
	if err := binary.Read(cryptrand.Reader, binary.LittleEndian, &seed); err != nil {
		log.Fatal(err)
	}
	rand.Seed(seed)
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
			str := []string{maxStr, minStr}
			return BetweenMinAndMax(str)
		}

		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(max - min) + min

		return ChampKV{Int: ret, Bint: true}, nil
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

		return ChampKV{Int: n, Bint: true}, nil
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

		resetSeed()
		// fmt.Println(time.Now().UnixNano())

		var ret int
		ret = rand.Intn(n+1)

		return ChampKV{Int: ret, Bint: true}, nil
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

		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(n)

		return ChampKV{Int: ret, Bint: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func GreaterThan(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		nStr := args[0]

		n, err := strconv.Atoi(nStr)
		if err != nil {
			return toReturn, err
		}

		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(int(math.MaxInt32) - n)
		ret += n

		return ChampKV{Int: ret, Bint: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func StrictlyGreaterThan(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		nStr := args[0]

		n, err := strconv.Atoi(nStr)
		if err != nil {
			return toReturn, err
		}

		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(int(math.MaxInt32) - n)
		ret += n + 1

		return ChampKV{Int: ret, Bint: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func EvenNumber(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 0 {
		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = (rand.Intn(int(math.MaxInt32))) * 2

		return ChampKV{Int: ret, Bint: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func OddNumber(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 0 {
		// rand.Seed(time.Now().UnixNano())
		var ret int
		ret = rand.Intn(int(math.MaxInt32))
		ret += (ret % 2) + 1

		return ChampKV{Int: ret, Bint: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}

func MultipleOf(args []string) (ChampKV, error) {
	var toReturn ChampKV

	if len(args) == 1 {
		var ret float64
		n, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return toReturn, errors.New("Impossible to parse parameter into float.")
		}

		// rand.Seed(time.Now().UnixNano())

		ret = rand.Float64() * n
		ret *= float64(rand.Intn(int(math.MaxFloat64/math.Abs(ret))))

		return ChampKV{Float: ret, Bfloat: true}, nil
	} else {
		return toReturn, errors.New("Bad argument number.")
	}
}