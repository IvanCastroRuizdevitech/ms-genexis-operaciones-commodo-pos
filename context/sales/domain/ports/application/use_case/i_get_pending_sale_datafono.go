package iusecase

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetPendingSaleDatafono interface {
    Execute(transactionId int) (*entities_main.Response[entities_sales.PendingSaleDatafono], error)
}

