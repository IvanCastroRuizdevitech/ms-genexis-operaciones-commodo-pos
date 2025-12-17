package entities

// InsertPreAuthorizationCustomerRequest maps to the JSONB payload expected by
// fnc_insertar_pre_autorizacion_cliente.
type InsertPreAuthorizationCustomerRequest struct {
	Codigo                 string `json:"codigo" binding:"required"`
	Surtidor               int    `json:"surtidor" binding:"required"`
	Cara                   int    `json:"cara" binding:"required"`
	Grado                  int    `json:"grado" binding:"required"`
	ProveedoresID          int    `json:"proveedores_id" binding:"required"`
	Preventa               *bool  `json:"preventa" binding:"required"`
	Estado                 string `json:"estado" binding:"required"`
	Usado                  string `json:"usado" binding:"required"`
	MetodoPago             int    `json:"metodo_pago" binding:"required"`
	MedioAutorizacion      string `json:"medio_autorizacion" binding:"required"`
	SerialDispositivo      string `json:"serial_dispositivo" binding:"required"`
	PromotorID             int    `json:"promotor_id" binding:"required"`
	TipoTransaccion        int    `json:"tipo_transaccion" binding:"required"`
	EstadoTransaccion      int    `json:"estado_transaccion" binding:"required"`
	TransaccionesOrigenID  int    `json:"transacciones_origen_id" binding:"required"`
	Autorizacion           string `json:"autorizacion" binding:"required"`
}

// InsertPreAuthorizationCustomerResult wraps the boolean result returned by
// fnc_insertar_pre_autorizacion_cliente.
type InsertPreAuthorizationCustomerResult struct {
	Success bool `json:"success"`
}
