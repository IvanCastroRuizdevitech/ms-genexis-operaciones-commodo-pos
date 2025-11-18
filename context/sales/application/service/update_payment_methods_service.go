package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdatePaymentMethodsClient struct {
    UseCase iusecase.IUpdatePaymentMethods
}

func (s *UpdatePaymentMethodsClient) Execute(request *entities_sales.UpdatePaymentMethodsRequest) (*entities_main.Response[entities_sales.UpdatePaymentMethodsResult], error) {
    return s.UseCase.Execute(request)
}

