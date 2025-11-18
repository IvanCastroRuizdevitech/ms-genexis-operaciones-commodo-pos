package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdatePaymentMethodsRepository interface {
    Update(request *entities_sales.UpdatePaymentMethodsRequest) (*entities_main.Response[entities_sales.UpdatePaymentMethodsResult], error)
}

