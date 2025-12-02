package entities

// ShiftSummary represents the aggregated sales summary per shift.
type ShiftSummary []map[string]any

type ShiftSummaryRequest struct {
	Pos         int    `json:"pos" binding:"required"`
	FechaInicio string `json:"fecha_inicio" binding:"required"`
	FechaFin    string `json:"fecha_fin" binding:"required"`
}
