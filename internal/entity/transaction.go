package entity

type Detail struct {
	Periodo               string  `json:"periodo"`
	TipoEntidad           string  `json:"tipoEntidad"`
	Entidad               string  `json:"entidad"`
	Region                string  `json:"region"`
	Provincia             string  `json:"provincia"`
	Persona               string  `json:"persona"`
	Genero                string  `json:"genero"`
	TipoCliente           string  `json:"tipoCliente"`
	InstrumentoCaptacion  string  `json:"instrumentoCaptacion"`
	Divisa                string  `json:"divisa"`
	CantidadInstrumento   int     `json:"cantidadInstrumento"`
	Balance               float64 `json:"balance"`
	TasaPromedioPonderado float64 `json:"tasaPromedioPonderado"`
}
