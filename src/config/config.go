package config

func GetDSN() string {
    var User = "root" // exemple : "root"
    var Password = "" // laisser à vide si pas de mot de passe
    var DatabaseName = ""
    var Charset = "utf8" // exemple : "utf8"

    var ret = User + ":" + Password + "@/" + DatabaseName + "?charset=" + Charset

    return ret
}

var Port = ":8000" // exemple : ":8000"