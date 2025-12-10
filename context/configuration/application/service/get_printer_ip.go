package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetPrinterIPClient struct {
	UseCase iusecase.IGetPrinterIP
}

func (s *GetPrinterIPClient) Execute() (*entities_main.Response[entities.PrinterIPResult], error) {
	return s.UseCase.Execute()
}
