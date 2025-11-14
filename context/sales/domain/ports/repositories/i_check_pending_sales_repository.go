package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICheckPendingSalesRepository interface {
    Check(request *entities_sales.CheckPendingSalesRequest) (*entities_main.Response[[]entities_sales.PendingSale], error)
}

