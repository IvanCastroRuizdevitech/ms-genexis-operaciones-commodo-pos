package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetShiftSummaryClient struct {
	GetShiftSummary iusecase.IGetShiftSummary
}

func (s *GetShiftSummaryClient) Execute(pos int, fechaInicio string, fechaFin string) (*entities_main.Response[entities_reports.ShiftSummary], error) {
	return s.GetShiftSummary.Execute(pos, fechaInicio, fechaFin)
}
