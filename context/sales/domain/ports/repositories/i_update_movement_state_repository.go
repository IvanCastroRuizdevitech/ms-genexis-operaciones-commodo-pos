package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateMovementStateRepository interface {
    Update(movementId int, request *entities_sales.UpdateMovementStateRequest) (*entities_main.Response[entities_sales.UpdateMovementStateResult], error)
}

