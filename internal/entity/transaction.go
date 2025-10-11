package entity

type Detail struct {
	Periodo               string  `json:"periodo"`
	TipoEntidad           string  `json:"tipoentidad"`
	Entidad               string  `json:"entidad"`
	Region                string  `json:"region"`
	Provincia             string  `json:"provincia"`
	Persona               string  `json:"persona"`
	Genero                string  `json:"genero"`
	TipoCliente           string  `json:"tipocliente"`
	InstrumentoCaptacion  string  `json:"instrumentocaptacion"`
	Divisa                string  `json:"divisa"`
	CantidadInstrumento   int     `json:"cantidadinstrumento"`
	Balance               float64 `json:"balance"`
	TasaPromedioPonderado float64 `json:"tasapromedioponderado"`
}
