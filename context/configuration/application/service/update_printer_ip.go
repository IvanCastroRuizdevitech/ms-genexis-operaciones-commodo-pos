package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdatePrinterIPClient struct {
	UseCase iusecase.IUpdatePrinterIP
}

func (s *UpdatePrinterIPClient) Execute(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error) {
	return s.UseCase.Execute(req)
}
