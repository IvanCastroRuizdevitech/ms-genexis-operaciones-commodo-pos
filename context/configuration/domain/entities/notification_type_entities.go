package entities

// NotificationTypesResponse models the JSON returned by
// public.fnc_obtener_tipo_notificacion().
type NotificationTypesResponse struct {
	Success bool                     `json:"success"`
	Total   int                      `json:"total"`
	Data    []map[string]interface{} `json:"data"`
	Error   string                   `json:"error,omitempty"`
}
