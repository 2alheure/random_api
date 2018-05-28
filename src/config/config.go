package config

/* Database config */
var User = "root" // exemple : "root"	REQUIS
var Password = "" // laisser à vide si pas de mot de passe
var DatabaseName = ""	// exemple : "base"		REQUIS
var Charset = "utf8" // exemple : "utf8"

func DSN() string {
	var ret = User + ":" + Password + "@/" + DatabaseName + "?charset=" + Charset
	return ret
}

/* Server config */
var Port = ":8000" // exemple : ":8000"
