package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckDatafonoCancellationsInProgressClient struct {
    UseCase iusecase.ICheckDatafonoCancellationsInProgress
}

func (s *CheckDatafonoCancellationsInProgressClient) Execute(request *entities_sales.DatafonoCancellationsInProgressRequest) (*entities_main.Response[entities_sales.DatafonoCancellationsInProgress], error) {
    return s.UseCase.Execute(request)
}

