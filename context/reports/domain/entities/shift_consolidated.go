package entities

// ShiftConsolidated represents the consolidated sales by shift across a date range.
type ShiftConsolidated []map[string]any

type ShiftConsolidatedRequest struct {
	FechaInicio string `json:"fecha_inicio" binding:"required"`
	FechaFin    string `json:"fecha_fin" binding:"required"`
}
