package config

/* Database config */
var User = "root" // exemple : "root"	REQUIS
var Password = "" // laisser à vide si pas de mot de passe
var Host = "localhost"	// exemple : "localhost"	REQUIS
var DatabaseName = ""	// exemple : "base"		REQUIS
var Charset = "utf8" // exemple : "utf8"

func DSN() string {
	var ret = User + ":" + Password + "@" + Host + "/" + DatabaseName + "?charset=" + Charset
	return ret
}

/* Server config */
var Port = ":8000" // exemple : ":8000"
