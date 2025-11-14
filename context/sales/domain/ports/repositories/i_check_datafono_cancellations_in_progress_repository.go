package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ICheckDatafonoCancellationsInProgressRepository interface {
    Check(request *entities_sales.DatafonoCancellationsInProgressRequest) (*entities_main.Response[entities_sales.DatafonoCancellationsInProgress], error)
}

