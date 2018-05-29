package generator

type RessourceKV struct {
	Champs		[]ChampKV		`json:",omitempty"`
}

type ChampKV struct {					// On met la clef du champ dedans ?
	Clef	string
	Int		int
	Float	float64
	String	string
	Bool	bool
}

