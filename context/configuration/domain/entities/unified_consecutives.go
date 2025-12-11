package entities

// UnifiedConsecutivesResponse mirrors the JSON returned by fnc_obtener_consecutivos_unificados.
type UnifiedConsecutivesResponse struct {
	Success bool                 `json:"success"`
	Total   int                  `json:"total"`
	Data    []UnifiedConsecutive `json:"data"`
}

// UnifiedConsecutive represents a single consecutive entry from COM or TIEN destinations.
type UnifiedConsecutive struct {
	ID                 int64   `json:"id"`
	TipoDocumento      string  `json:"tipo_documento"`
	Prefijo            *string `json:"prefijo"`
	FechaInicio        *string `json:"fecha_inicio"`
	FechaFin           *string `json:"fecha_fin"`
	ConsecutivoInicial *int64  `json:"consecutivo_inicial"`
	ConsecutivoFinal   *int64  `json:"consecutivo_final"`
	ConsecutivoActual  *int64  `json:"consecutivo_actual"`
	Estado             *string `json:"estado"`
	Destino            string  `json:"destino"`
}
