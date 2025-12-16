package entities

import "encoding/json"

// DispenserDetailsFunctionResponse mirrors the JSON object returned by
// fnc_obtener_surtidores_detalles(), keeping the inner data untyped to avoid
// coupling to DB column names.
type DispenserDetailsFunctionResponse struct {
	Success bool            `json:"success"`
	Total   int             `json:"total"`
	Data    json.RawMessage `json:"data"`
}

// DispenserDetailsByFamiliesRequest wraps the families filter for
// fnc_obtener_surtidores_detalles_por_familias().
type DispenserDetailsByFamiliesRequest struct {
	FamiliesIDs []int32 `json:"families_ids" binding:"required"`
}
