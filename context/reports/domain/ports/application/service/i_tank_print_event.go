package iservice

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ITankPrintEvent interface {
	Execute(tankIDs []int) (*entities_main.Response[entities_reports.TankPrintEventResult], error)
}
