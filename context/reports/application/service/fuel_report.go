package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type FuelReportClient struct {
	GetFuelReport iusecase.IGetFuelReport
}

func (s *FuelReportClient) Execute(fecha string) (*entities_main.Response[entities_reports.FuelReport], error) {
	return s.GetFuelReport.Execute(fecha)
}
