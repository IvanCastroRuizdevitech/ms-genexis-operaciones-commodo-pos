package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type FuelEntryReportClient struct {
    UseCase iusecase.IFuelEntryReport
}

func (s *FuelEntryReportClient) Execute(request *entities_sales.FuelEntryReportRequest) (*entities_main.Response[entities_sales.FuelEntryReportResult], error) {
    return s.UseCase.Execute(request)
}

