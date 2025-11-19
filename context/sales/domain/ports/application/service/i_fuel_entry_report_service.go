package iservice

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IFuelEntryReport interface {
    Execute(request *entities_sales.FuelEntryReportRequest) (*entities_main.Response[entities_sales.FuelEntryReportResult], error)
}

