package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICheckReadySalesRepository interface {
    Check(request *entities_sales.CheckReadySalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error)
}

