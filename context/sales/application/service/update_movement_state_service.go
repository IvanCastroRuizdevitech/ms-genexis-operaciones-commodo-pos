package service

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    iusecase "ms-genexis-pos-operaciones/context/sales/domain/ports/application/use_case"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type UpdateMovementStateClient struct {
    UseCase iusecase.IUpdateMovementState
}

func (s *UpdateMovementStateClient) Execute(movementId int, request *entities_sales.UpdateMovementStateRequest) (*entities_main.Response[entities_sales.UpdateMovementStateResult], error) {
    return s.UseCase.Execute(movementId, request)
}

