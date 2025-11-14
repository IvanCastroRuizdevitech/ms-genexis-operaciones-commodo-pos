package iusecase

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICheckReadySales interface {
    Execute(request *entities_sales.CheckReadySalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error)
}

