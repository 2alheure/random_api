package generator;

import (
	"fmt"
	"errors"
	
	"github.com/lucasjones/reggen"
)

func TestStr() {
	fmt.Println("test str_gen");
}

func FromRegex(args []string) (ChampKV, error) {
	var toReturn ChampKV
	var errReturn error

	if len(args) == 1 {
		regex := args[0]
		
		str, err := reggen.Generate(regex, 10)
		if err != nil {
			errReturn = errors.New(fmt.Sprintf("Unable to generate from regex \"%s\".", regex))
		} else {
			return ChampKV{String: str, Bstring: true}, nil
		}
	} else { 
		errReturn = errors.New(fmt.Sprintf("Bad argument number: Expected 1, Received %d.", len(args)))
	}
	return toReturn, errReturn
}