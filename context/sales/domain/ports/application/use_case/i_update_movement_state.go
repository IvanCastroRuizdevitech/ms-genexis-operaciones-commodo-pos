package iusecase

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateMovementState interface {
    Execute(movementId int, request *entities_sales.UpdateMovementStateRequest) (*entities_main.Response[entities_sales.UpdateMovementStateResult], error)
}

