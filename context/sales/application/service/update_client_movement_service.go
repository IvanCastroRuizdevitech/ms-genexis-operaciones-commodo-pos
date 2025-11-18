package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateClientMovementClient struct {
    UseCase iusecase.IUpdateClientMovement
}

func (s *UpdateClientMovementClient) Execute(request *entities_sales.UpdateClientMovementRequest) (*entities_main.Response[entities_sales.UpdateClientMovementResult], error) {
    return s.UseCase.Execute(request)
}

