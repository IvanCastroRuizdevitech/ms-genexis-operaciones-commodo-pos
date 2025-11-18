package iusecase

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetFuelReport interface {
	Execute(fecha string) (*entities_main.Response[entities_reports.FuelReport], error)
}
