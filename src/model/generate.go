package model

import (
	"fmt"
	"math/rand"
	"time"
	generator "../generator"
)

type generatorFunction func([]string)(generator.ChampKV, error)

type reducer struct {
	Clef		string
	Function	generatorFunction
	Params		[]string
}

var RuleSet = map[int]generatorFunction{		// Sert à récupérer toutes les fonctions des règles par leur id
	1: generator.FromRegex,
	2: generator.StrictlyLowerThan,
	3: generator.StrictlyGreaterThan,
	4: generator.LowerThan,
	5: generator.GreaterThan,
	6: generator.Equal,
	// 7: generator.EvenNumber,
	// 8: generator.OddNumber,
	9: generator.MultipleOf,
	// 10: generator.Dictionnary,
	11: generator.BetweenMinAndMax,
}

func Generate(ressource_id, nombre int) ([]generator.RessourceKV, int) {
	rand.Seed(time.Now().UnixNano())
	ret := make([]generator.RessourceKV, nombre)

	reduc, err_code, err := GetReducer(ressource_id)
	n := len(err)
	if n > 0 {
		return ret, err_code
	}

	var errorReport error

	for i := 0; i<nombre; i++ {

		var ress = generator.RessourceKV{Champs: make([]generator.ChampKV, len(reduc))}
		var field generator.ChampKV

		for j, red := range reduc {

			field, errorReport = red.Function(red.Params)
			if errorReport != nil {
				return ret, 500
			}

			field.Clef = red.Clef

			ress.Champs[j] = field
		}
		
		ret[i] = ress
	}

	return ret, 200
}

func GetReducer(ressource_id int) (ret []reducer, errCode int, messages []string) {
	ret = make([]reducer, 0)
	errCode = 200
	messages = make([]string, 0)
	
	ressource, err := GetRessource(ressource_id)
	if err != nil {
		errCode = 404
		errare := fmt.Errorf("Ressource d'id %d introuvable.", ressource_id)
		messages = append(messages, errare.Error())
		return
	}

	ressource.Hydrate()



	for _, champ := range ressource.Champs {
		regle_id:= champ.Regle.Id
		if _, ok := RuleSet[regle_id]; !ok {
			errCode = 409
			errare := fmt.Errorf("La fonction permettant de traiter la règle `%s` (regle_id: %d) n'est pas encore implémentée.", champ.Regle.Nom, regle_id)
			messages = append(messages, errare.Error())
			continue
		}

		if !champ.Clef.Valid || champ.Regle == nil  || regle_id == 0  {
			errCode = 409
			errare := fmt.Errorf("Impossible d'hydrater correctement le champ `%s` ou sa règle `%s`.", champ.Clef.Value, champ.Regle.Nom)
			messages = append(messages, errare.Error())
			continue
		}

		params := champ.Regle.Parametres
		var givableParams []string
		
		if len(params) == 0 {			// !!! Attention il y a des règles sans paramètre !!!
			errCode = 409
			errare := fmt.Errorf("Aucun paramètre pour la règle `%s` qui en attendait.", champ.Regle.Nom)
			messages = append(messages, errare.Error())
			continue
		} else {

			for _, param := range params {

				if param.Type == "" || param.Value == "" {
					errCode = 409
					errare := fmt.Errorf("Paramètre `%s` (type %s) invalide pour le champ `%s`.", param.Value, param.Type, champ.Clef.Value)
					messages = append(messages, errare.Error())
					continue
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

	return
}