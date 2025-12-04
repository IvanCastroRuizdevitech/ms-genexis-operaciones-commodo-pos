package service

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type SetInvoiceAttributesService struct {
	UseCase iusecase.ISetInvoiceAttributes
}

func (s *SetInvoiceAttributesService) Execute(request *entities_sales.SetInvoiceAttributesRequest) (*entities_main.Response[entities_sales.SetInvoiceAttributesResult], error) {
	return s.UseCase.Execute(request)
}
