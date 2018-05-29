package config

func DSN() string {
    var User = "randomapi" // exemple : "root"
    var Password = "tjJR443an5ftHu8U" // laisser à vide si pas de mot de passe
    var DatabaseName = "randomapi"
    var Charset = "utf8_bin" // exemple : "utf8"

    var ret = User + ":" + Password + "@/" + DatabaseName + "?charset=" + Charset

    return ret
}

var Port = ":7777" // exemple : ":8000"