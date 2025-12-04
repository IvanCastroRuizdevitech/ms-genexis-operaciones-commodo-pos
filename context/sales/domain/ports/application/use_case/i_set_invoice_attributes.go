package iusecase

import (
	entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
	entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ISetInvoiceAttributes interface {
	Execute(request *entities_sales.SetInvoiceAttributesRequest) (*entities_main.Response[entities_sales.SetInvoiceAttributesResult], error)
}
