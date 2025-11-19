package usecase

import (
    "log"
    entities_sales "ms-genexis-pos-operaciones/context/sales/domain/entities"
    irepo "ms-genexis-pos-operaciones/context/sales/domain/ports/repositories"
    entities_main "ms-genexis-pos-operaciones/domain/entities"
)

type ReprintSale struct {
    Repository irepo.IReprintSaleRepository
}

func (u *ReprintSale) Execute(movementId int) (*entities_main.Response[entities_sales.ReprintSaleResult], error) {
    result, err := u.Repository.Reprint(movementId)
    if err != nil {
        log.Println("ReprintSale usecase error:", err)
        return nil, err
    }
    return result, nil
}

