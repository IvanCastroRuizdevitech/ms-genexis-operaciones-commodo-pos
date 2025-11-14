package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckReadySalesClient struct {
    UseCase iusecase.ICheckReadySales
}

func (s *CheckReadySalesClient) Execute(request *entities_sales.CheckReadySalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error) {
    return s.UseCase.Execute(request)
}

