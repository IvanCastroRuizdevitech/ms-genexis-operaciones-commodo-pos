package service

import (
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
)

type DayClosingReportClient struct {
    GetDayClosingReport iusecase.IGetDayClosingReport
}

func (s *DayClosingReportClient) Execute(fecha string) (*entities_main.Response[entities_reports.DayClosingReport], error) {
    return s.GetDayClosingReport.Execute(fecha)
}

