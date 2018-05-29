package generator;

import (
	"fmt"
	"strconv"
	
	"github.com/lucasjones/reggen"

	help "../helper"
);

func TestStr() {
	fmt.Println("test str_gen");
}

func FromRegex(args []string) (ChampKV, error) {
	var toReturn ChampKV
	var errReturn error

	if len(args) == 1 {
		regex := args[0]
		
		var err error
		toReturn.String, err = reggen.NewGenerator(regex)
		if err != nil {
			errReturn.New("Unable to generate from regex.")
		} 
	} else { 
		errReturn.New("Bad Argument Number.")
	}
	return toReturn, errReturn
}