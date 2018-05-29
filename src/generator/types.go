package generator

type RessourceKV struct {
	Champs		[]ChampKV		`json:"-,omitempty"`
}

type ChampKV struct {					// On met la clef du champ dedans ?
	Clef	string			`json:"clef,omitempty"`
	Int		int				`json:"valeuri,omitempty"`
	Float	float64			`json:"valeurf,omitempty"`
	String	string			`json:"valeurs,omitempty"`
	Bool	bool			`json:"valeurb,omitempty"`
}

