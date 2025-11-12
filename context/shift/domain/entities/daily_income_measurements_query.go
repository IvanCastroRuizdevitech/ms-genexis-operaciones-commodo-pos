package entities

import "time"

type DailyIncomeMeasurementsQuery struct {
    FechaInicio time.Time `query:"fecha_inicio" validate:"required,datetime"`
    FechaFin    time.Time `query:"fecha_fin" validate:"required,datetime"`
}

