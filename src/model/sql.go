package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	help "random_api/src/helper"
)

var Bdd *sql.DB

func TestSql() {
	fmt.Println("test sql")
}

func InitBdd() *sql.DB {
	var err error
	Bdd, err = sql.Open("mysql", "root:@/alea_data_est?charset=utf8")
	help.CheckErr(err)

	return Bdd
}

// func getRessourceFullStructure(id int) string {
// 	str := "
	// SELECT 
	// 	champ.clef AS clef,
	// 	champ.id AS champ_id,
	// 	regle.nom AS regle,
	// 	regle.id AS regle_id,
	// 	GROUP_CONCAT(
	// 		CONCAT(
	// 			'{\"id\": \"', champ_parametre.id, '\", '
	// 			'\"type\": \"', parametre.nom, '\", '
	// 			'\"value\": \"', champ_parametre.valeur, '\"}'
	// 		) ORDER BY regle_parametre.id
	// 	) AS parametres
	// FROM ressource
	// 	LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id
	// 	LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id
	// 	LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id
	// 	LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id
	// 	LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id
	// WHERE ressource.id = 1
	// GROUP BY champ_parametre.champ_id
	// ORDER BY clef
// 	"

// 	return str
// }

// func getRessourceStructure(id int) string {
// 	str := "
// 	SELECT 
// 		ressource.*,
// 		champ.clef
// 	FROM ressource
// 		LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id
// 	WHERE ressource.id = ?
// 	ORDER BY clef
// 	"

// 	return str
// }

