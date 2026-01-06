package entities

// UpdateComandaStatusRequest maps to restaurante.actualizar_estado_comanda.
type UpdateComandaStatusRequest struct {
	ComandaID     int64 `json:"p_comanda_id" binding:"required"`
	NuevoEstadoID int16 `json:"p_nuevo_estado_id" binding:"required"`
}

// UpdateComandaStatusResult mirrors the json_build_object returned by actualizar_estado_comanda.
type UpdateComandaStatusResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Status  int    `json:"status,omitempty"`
}
