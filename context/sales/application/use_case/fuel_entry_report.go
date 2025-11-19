package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type FuelEntryReport struct {
    Repository irepo.IFuelEntryReportRepository
}

func (u *FuelEntryReport) Execute(request *entities_sales.FuelEntryReportRequest) (*entities_main.Response[entities_sales.FuelEntryReportResult], error) {
    result, err := u.Repository.Generate(request)
    if err != nil {
        log.Println("FuelEntryReport usecase error:", err)
        return nil, err
    }
    return result, nil
}

