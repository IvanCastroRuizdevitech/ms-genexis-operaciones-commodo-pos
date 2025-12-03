package irepositories

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetShiftConsolidatedRepository interface {
	Get(fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftConsolidated], error)
}
