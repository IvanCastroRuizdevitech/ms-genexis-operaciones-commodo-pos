package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type AssignCustomerDataClient struct {
    UseCase iusecase.IAssignCustomerData
}

func (s *AssignCustomerDataClient) Execute(request *entities_sales.AssignCustomerDataRequest) (*entities_main.Response[entities_sales.AssignCustomerDataResult], error) {
    return s.UseCase.Execute(request)
}

