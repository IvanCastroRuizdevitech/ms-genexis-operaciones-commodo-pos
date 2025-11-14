package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type CheckPendingSalesClient struct {
    UseCase iusecase.ICheckPendingSales
}

func (s *CheckPendingSalesClient) Execute(request *entities_sales.CheckPendingSalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error) {
    return s.UseCase.Execute(request)
}

