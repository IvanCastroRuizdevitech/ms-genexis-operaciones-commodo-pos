package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type TankPrintEventClient struct {
	CreateTankPrintEvent iusecase.ICreateTankPrintEvent
}

func (s *TankPrintEventClient) Execute(tankIDs []int) (*entities_main.Response[entities_reports.TankPrintEventResult], error) {
	return s.CreateTankPrintEvent.Execute(tankIDs)
}
