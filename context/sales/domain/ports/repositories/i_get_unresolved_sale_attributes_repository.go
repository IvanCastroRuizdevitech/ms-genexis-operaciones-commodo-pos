package irepositories

import (
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type IGetUnresolvedSaleAttributesRepository interface {
    Get(movementId int) (*entities_main.Response[entities_sales.UnresolvedSaleAttributes], error)
}

