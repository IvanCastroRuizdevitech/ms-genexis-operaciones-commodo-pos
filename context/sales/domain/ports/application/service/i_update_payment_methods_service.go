package iservice

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdatePaymentMethods interface {
    Execute(request *entities_sales.UpdatePaymentMethodsRequest) (*entities_main.Response[entities_sales.UpdatePaymentMethodsResult], error)
}

