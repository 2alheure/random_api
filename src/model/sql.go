package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	help "../helper"
	config "../config"
)

var Bdd *sql.DB

func TestSql() {
	fmt.Println("test sql")
}

func InitBdd() {
	var err error
	Bdd, err = sql.Open("mysql", config.DSN())
	help.CheckErr(err)
}

/** Liste des requêtes :

Récupérer une ressource entière (pour Hydrater)
	SELECT 
	champ.clef AS clef,
	champ.id AS champ_id,
	regle.nom AS regle,
	regle.id AS regle_id,
	CONCAT(
		'[',
		GROUP_CONCAT(
			CONCAT(
				'{\"id\": ', champ_parametre.id, ', '
				'\"type\": \"', parametre.nom, '\", '
				'\"value\": \"', champ_parametre.valeur, '\"}'
			) ORDER BY regle_parametre.id
		),
		']'
	) AS parametres
	FROM ressource
	LEFT OUTER JOIN champ ON champ.ressource_id = ressource.id
	LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id
	LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id
	LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id
	LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id
	WHERE ressource.id = ?
	GROUP BY champ_parametre.champ_id
	ORDER BY clef


Récupérer un champ entier :
	SELECT 
		champ.clef AS clef,
		champ.id AS champ_id,
		regle.nom AS regle,
		regle.id AS regle_id,
		CONCAT(
			'[',
			GROUP_CONCAT(
				CONCAT(
					'{\"id\": ', champ_parametre.id, ', '
					'\"type\": \"', parametre.nom, '\", '
					'\"value\": \"', champ_parametre.valeur, '\"}'
				) ORDER BY regle_parametre.id
			),
			']'
		) AS parametres
	FROM champ
	LEFT OUTER JOIN champ_parametre ON champ_parametre.champ_id = champ.id
	LEFT OUTER JOIN regle_parametre ON champ_parametre.regle_parametre_id = regle_parametre.id
	LEFT OUTER JOIN regle ON regle_parametre.regle_id = regle.id
	LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id
	WHERE champ.id = ?
	GROUP BY champ_parametre.champ_id


Récupérer l'ensemble des règles et leurs paramètres :
	SELECT
    regle.nom AS regle,
	regle.id AS regle_id,
	CONCAT(
		'[',
		GROUP_CONCAT(
			CONCAT(
				'{\"id\": ', parametre.id, ', '
				'\"type\": \"', parametre.nom, '\"}'
			) ORDER BY regle_parametre.id
		),
		']'
	) AS parametres
	FROM regle
	LEFT OUTER JOIN regle_parametre ON regle.id = regle_parametre.regle_id
	LEFT OUTER JOIN parametre ON regle_parametre.parametre_id = parametre.id
	GROUP BY regle_parametre.regle_id
*/