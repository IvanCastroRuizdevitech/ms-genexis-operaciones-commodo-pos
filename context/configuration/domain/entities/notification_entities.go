package entities

// ProcessNotificationRequest represents the payload to invoke
// public.prc_procesar_notificacion.
type ProcessNotificationRequest struct {
	TipoNotificacion int64  `json:"tipo_notificacion" binding:"required"`
	Data             string `json:"data" binding:"required"`
	Prioridad        bool   `json:"prioridad" binding:"required"`
}

// ProcessNotificationResult captures the JSON response returned by
// public.prc_procesar_notificacion through its INOUT parameter.
type ProcessNotificationResult struct {
	CodigoRespuesta int    `json:"codigo_respuesta"`
	Estado          string `json:"estado"`
	IdNotificacion  int64  `json:"id_notificacion"`
}
