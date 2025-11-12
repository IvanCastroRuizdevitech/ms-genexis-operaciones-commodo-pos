package entities

type OpeningShiftHttp struct {
	PersonasId      int         `json:"personas_id"`
	Surtidores      []int32     `json:"surtidores"`
	FechaInicio     string      `json:"fecha_inicio"`
	EquiposId       int32       `json:"equipos_id"`
	EmpresasId      int32       `json:"empresas_id"`
	Atributos       Attributes  `json:"atributos"`
	AjustePeriodico interface{} `json:"ajustePeriodico"`
}

type OpeningShiftHttpResponse struct {
	Estado       int          `json:"status"`
	Turno        int          `json:"turno,omitempty"`
	Respuesta    *interface{} `json:"respuesta,omitempty"`
	Mensaje      string       `json:"mensaje,omitempty"`
	Data         interface{}  `json:"data,omitempty"`
	FechaProceso string       `json:"fechaProceso,omitempty"`
	CodigoError  string       `json:"codigoError,omitempty"`
	TipoError    string       `json:"tipoError,omitempty"` // Omitido si está vacío
}
