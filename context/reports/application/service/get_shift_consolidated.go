package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetShiftConsolidatedClient struct {
	GetShiftConsolidated iusecase.IGetShiftConsolidated
}

func (s *GetShiftConsolidatedClient) Execute(fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftConsolidated], error) {
	return s.GetShiftConsolidated.Execute(fechaInicio, fechaFin)
}
