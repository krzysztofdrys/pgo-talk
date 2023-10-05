package city

type City struct {
	Type       string `json:"type"`
	Properties struct {
		OID      int         `json:"OID"`
		UCUID    int         `json:"UCUID"`
		Name     string      `json:"Name"`
		GENC0    string      `json:"GENC0"`
		Ctry     string      `json:"Ctry"`
		GENC1    string      `json:"GENC1"`
		Prvn     string      `json:"Prvn"`
		Pop      int         `json:"Pop"`
		PopCls   int         `json:"PopCls"`
		Lat      float64     `json:"Lat"`
		Lon      float64     `json:"Lon"`
		PopLvl   int         `json:"PopLvl"`
		UNcityCd interface{} `json:"UNcityCd"`
	} `json:"properties"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

type Dataset struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Crs  struct {
		Type       string `json:"type"`
		Properties struct {
			Name string `json:"name"`
		} `json:"properties"`
	} `json:"crs"`
	Features []City `json:"features"`
}
