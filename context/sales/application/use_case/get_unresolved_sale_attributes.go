package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type GetUnresolvedSaleAttributes struct {
    Repository irepo.IGetUnresolvedSaleAttributesRepository
}

func (u *GetUnresolvedSaleAttributes) Execute(movementId int) (*entities_main.Response[entities_sales.UnresolvedSaleAttributes], error) {
    result, err := u.Repository.Get(movementId)
    if err != nil {
        log.Println("GetUnresolvedSaleAttributes usecase error:", err)
        return nil, err
    }
    return result, nil
}

