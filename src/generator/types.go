package generator

import (
	"fmt"
	"strconv"
	"bytes"
)

type RessourceKV struct {
	Champs		[]ChampKV		`json:"omitempty"`
}

type ChampKV struct {					// On met la clef du champ dedans ?
	Clef		string			`json:"clef,omitempty"`
	Int			int				`json:"valeuri,omitempty"`
	Float		float64			`json:"valeurf,omitempty"`
	String		string			`json:"valeurs,omitempty"`
	Bool		bool			`json:"valeurb,omitempty"`
	Bint		bool
	Bfloat		bool
	Bstring		bool
	Bbool		bool
}

func (ress RessourceKV) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")
	length := len(ress.Champs)
	i := 0

	for _, val := range ress.Champs {
		if val.Bint {
			buffer.WriteString(fmt.Sprintf("\"%s\":\"%s\"", val.Clef, strconv.Itoa(val.Int)))
		}
		if val.Bfloat {
			buffer.WriteString(fmt.Sprintf("\"%s\":\"%s\"", val.Clef, strconv.FormatFloat(val.Float, 'g', -1, 64)))
		}
		if val.Bstring {
			buffer.WriteString(fmt.Sprintf("\"%s\":\"%s\"", val.Clef, val.String))
		}
		if val.Bbool {
			buffer.WriteString(fmt.Sprintf("\"%s\":\"%s\"", val.Clef, strconv.FormatBool(val.Bool)))
		}

		i++
		if i < length {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString("}")

	return buffer.Bytes(), nil

}