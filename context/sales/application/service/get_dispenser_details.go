package service

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetDispenserDetailsClient struct {
	UseCase iusecase.IGetDispenserDetails
}

func (s *GetDispenserDetailsClient) Execute() (*entities_main.Response[[]entities_sales.DispenserDetail], error) {
	return s.UseCase.Execute()
}
