package entities

type FuelPumpsRequest struct {
    TurnoId    int64 `json:"turno_id" binding:"required"`
    EquiposId int64 `json:"equipos_id" binding:"required"`
}
