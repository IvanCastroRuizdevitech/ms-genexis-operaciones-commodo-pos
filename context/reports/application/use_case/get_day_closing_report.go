package usecase

import (
    "log"
    irepositories "ms-genexis-pos-operaciones/context/reports/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
    entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
)

type GetDayClosingReport struct {
    Repository irepositories.IGetDayClosingReportRepository
}

func (u *GetDayClosingReport) Execute(fecha string) (*entities_main.Response[entities_reports.DayClosingReport], error) {
    result, err := u.Repository.GetReport(fecha)
    if err != nil {
        log.Println("GetDayClosingReport error:", err)
        return nil, err
    }
    return result, nil
}

