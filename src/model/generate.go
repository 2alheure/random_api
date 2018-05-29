package model

import (
	_ "fmt"
	"errors"
	generator "random_api/src/generator"
)

type generatorFunction func([]string)(generator.ChampKV, error)

type reducer struct {
	Clef		string					// string ou generator.ChampKV, voir comment ça évolue
	Function	generatorFunction
	Params		[]string
}

var RuleSet = map[int]generatorFunction{		// Sert à récupérer toutes les fonctions des règles par leur id
	// 1: generator.FromRegex,
	2: generator.StrictlyLowerThan,
	// 3: generator.StrictlyGreaterThan,
	4: generator.LowerThan,
	// 5: generator.GreaterThan,
	6: generator.Equal,
	7: generator.EvenNumber,
	8: generator.OddNumber,
	// 9: generator.MultipleOf,
	// 10: generator.Dictionnary,
	11: generator.BetweenMinAndMax,
}

func Generate(ressource_id, nombre int) ([]generator.RessourceKV, int) {
	ret := make([]generator.RessourceKV, nombre)

	reduc, err_code, err := GetReducer(ressource_id)
	if err != nil {
		return ret, err_code
	}

	var ress = generator.RessourceKV{Champs: make([]generator.ChampKV, len(reduc))}
	var errorReport error
	for i := 0; i<nombre; i++ {
		var field generator.ChampKV

		for i, red := range reduc {

			field, errorReport = red.Function(red.Params)
			if errorReport != nil {
				return ret, 500
			}

			field.Clef = red.Clef

			ress.Champs[i] = field
		}

		ret[nombre] = ress
	}

	return ret, 200
}

func GetReducer(ressource_id int) ([]reducer, int, error) {
	var ret []reducer
	ressource, err := GetRessource(ressource_id)
	if err != nil {
		return ret, 404, errors.New("Ressource irrécupérable.")
	}

	ressource.Hydrate()



	for _, champ := range ressource.Champs {
		regle_id:= champ.Regle.Id

		if !champ.Clef.Valid || champ.Regle == nil  || regle_id == 0  {
			return ret, 409, errors.New("Problème avec la règle.")
		}

		params := champ.Regle.Parametres
		var givableParams []string
		
		if len(params) == 0 {			// !!! Attention il y a des règles sans paramètre !!!
			return ret, 409, errors.New("Aucun paramètre pour la regle.")
		} else {

			for _, param := range params {

				if param.Type == "" || param.Value == "" {
					return ret, 409, errors.New("Paramètre invalide.")
				} else {
					givableParams = append(givableParams, param.Value)
				}
			}
		}

		ret = append(ret, reducer{
			champ.Clef.String,
			RuleSet[regle_id],
			givableParams,
		})
	}

	return ret, 0, nil
}