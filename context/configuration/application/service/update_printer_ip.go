package service

import (
	"ms-genexis-pos-operaciones/context/configuration/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/configuration/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdatePrinterIPClient struct {
	UpdateUseCase    iusecase.IUpdatePrinterIP
	PrintAfterUpdate iusecase.IPrintTestAfterUpdate
}

func (s *UpdatePrinterIPClient) Execute(req *entities.PrinterIPUpdateRequest) (*entities_main.Response[entities.PrinterIPUpdateResult], error) {
	resp, err := s.UpdateUseCase.Execute(req)
	if err != nil {
		return nil, err
	}

	if s.PrintAfterUpdate != nil {
		return s.PrintAfterUpdate.Execute(resp)
	}

	return resp, nil
}
