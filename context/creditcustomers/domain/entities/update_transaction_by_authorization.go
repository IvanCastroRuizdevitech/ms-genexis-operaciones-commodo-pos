package entities

// UpdateTransactionByAuthorizationRequest maps to public.fnc_actualizar_transaccion_por_autorizacion
// and carries all fields required to refresh a transaction identified by its authorization UUID.
type UpdateTransactionByAuthorizationRequest struct {
	Autorizacion                string         `json:"i_autorizacion" binding:"required"`                   // uuid
	Surtidor                    int            `json:"i_surtidor" binding:"required"`                       // integer
	Cara                        int            `json:"i_cara" binding:"required"`                           // integer
	Grado                       int            `json:"i_grado" binding:"required"`                          // integer
	DocumentoCliente            string         `json:"i_documento_cliente" binding:"required"`              // text
	PlacaVehiculo               string         `json:"i_placa_vehiculo" binding:"required"`                 // text
	MontoMaximo                 float64        `json:"i_monto_maximo" binding:"required"`                   // numeric
	CantidadMaxima              float64        `json:"i_cantidad_maxima" binding:"required"`                // numeric
	ClienteNombre               string         `json:"i_cliente_nombre" binding:"required"`                 // text
	VehiculoOdometro            string         `json:"i_vehiculo_odometro" binding:"required"`              // text
	Trama                       map[string]any `json:"i_trama" binding:"required"`                          // json
	EstadoTransaccion           int            `json:"i_estado_transaccion" binding:"required"`             // smallint
	DocumentoConductor          string         `json:"i_documento_conductor" binding:"required"`            // text
	ConductorNombre             string         `json:"i_conductor_nombre" binding:"required"`               // text
	ClienteTipoIdentificacionID int            `json:"i_cliente_tipo_identificacion_id" binding:"required"` // integer
}

// UpdateTransactionByAuthorizationResult mirrors the json_build_object returned by
// fnc_actualizar_transaccion_por_autorizacion.
type UpdateTransactionByAuthorizationResult struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
