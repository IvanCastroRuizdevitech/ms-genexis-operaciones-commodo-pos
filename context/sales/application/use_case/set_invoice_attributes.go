package usecase

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	irepositories "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type SetInvoiceAttributes struct {
	Repository irepositories.ISetInvoiceAttributesRepository
}

func (uc *SetInvoiceAttributes) Execute(request *entities_sales.SetInvoiceAttributesRequest) (*entities_main.Response[entities_sales.SetInvoiceAttributesResult], error) {
	return uc.Repository.Update(request)
}
