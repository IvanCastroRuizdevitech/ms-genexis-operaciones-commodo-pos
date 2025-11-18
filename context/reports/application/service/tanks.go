package service

import (
	entities_reports "ms-genexis-pos-operaciones/context/reports/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/reports/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type TanksClient struct {
	GetTanks iusecase.IGetTanks
}

func (s *TanksClient) Execute() (*entities_main.Response[[]entities_reports.TankBodega], error) {
	return s.GetTanks.Execute()
}
