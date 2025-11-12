package entities

type OpeningShiftRequest struct {
	Usuario string `json:"usuario" binding:"required"`
	Clave   string `json:"clave" binding:"required"`

	Surtidores      []int32     `json:"surtidores" binding:"required"`
	FechaInicio     string      `json:"fecha_inicio" binding:"required"`
	EquiposId       int32       `json:"equipos_id" binding:"required"`
	EmpresasId      int32       `json:"empresas_id" binding:"required"`
	Atributos       Attributes  `json:"atributos" binding:"required"`
	AjustePeriodico interface{} `json:"ajustePeriodico"`
}

type InitialTotals struct {
	Surtidor               int32   `json:"surtidor"`
	Cara                   int32   `json:"cara"`
	Manguera               int32   `json:"manguera"`
	Grado                  int32   `json:"grado"`
	Isla                   int32   `json:"isla"`
	ProductoIdentificador  int64   `json:"productoIdentificador"`
	ProductoDescripcion    string  `json:"productoDescripcion"`
	FamiliaIdentificador   int32   `json:"familiaIdentificador"`
	FamiliaDescripcion     string  `json:"familiaDescripcion"`
	AcumuladoVolumen       int64   `json:"acumuladoVolumen"`
	AcumuladoVolumenReal   int64   `json:"acumuladoVolumenReal"`
	AcumuladoVenta         int64   `json:"acumuladoVenta"`
	Factor_inventario      int64   `json:"factor_inventario"`
	Factor_volumen_parcial int64   `json:"factor_volumen_parcial"`
	Factor_importe_parcial int64   `json:"factor_importe_parcial"`
	Factor_precio          int64   `json:"factor_precio"`
	Precio                 float64 `json:"precio"`
}

type Attributes struct {
	Saldo                  int64           `json:"saldo"`
	TotalizadoresIniciales []InitialTotals `json:"totalizadoresIniciales"`
}
