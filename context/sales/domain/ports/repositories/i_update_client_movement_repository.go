package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IUpdateClientMovementRepository interface {
    Update(request *entities_sales.UpdateClientMovementRequest) (*entities_main.Response[entities_sales.UpdateClientMovementResult], error)
}

